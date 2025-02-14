<template>
    <div>
        <video autoplay controls="true" ref="userVideo"></video>
        <video autoplay controls="true" ref="partnerVideo"></video>
    </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const userVideo = ref(null);
const userStream = ref(null);
const partnerVideo = ref(null);
const peerRef = ref(null);
const websocketRef = ref(null);

const BASE_URL = `ws://localhost:8000/join?roomID=${route.params.roomID}`;

const openCamera = async () => {
  const allDevices = await navigator.mediaDevices.enumerateDevices();
  const cameras = allDevices.filter(
    (device) => device.kind == "videoinput"
  );
  console.log('cameras', cameras);
  
  const constraints = {
    audio: true,
    video: {
      // cursor: "always"
      // deviceId: cameras[1].deviceId,
    }
  };
  
  try {
    return await navigator.mediaDevices.getUserMedia(constraints);
  } catch (err) {
    console.error(err);
    return null;
  }
};

const callUser = () => {
  console.log("calling other user");
  peerRef.value = createPeer();
  
  userStream.value.getTracks().forEach((track) => {
    peerRef.value.addTrack(track, userStream.value);
  });
};

const createPeer = () => {
  console.log("creating peer connection");
  
  const peer = new RTCPeerConnection({
    iceServers: [{ urls: "stun:stun.l.google.com:19302" }]
  });
  
  peer.onnegotiationneeded = handleNegotiationNeeded;
  peer.onicecandidate = handleIceCandidateEvent;
  peer.ontrack = handleTrackEvent;
  
  return peer;
};

const handleNegotiationNeeded = async () => {
  console.log("creating offer");
  try {
    const myOffer = await peerRef.value.createOffer();
    await peerRef.value.setLocalDescription(myOffer);
    
    websocketRef.value.send(JSON.stringify({ offer: peerRef.value.localDescription }));
  } catch (err) {
    console.error("Error creating offer:", err);
  }
};

const handleIceCandidateEvent = (e) => {
  console.log("found ice candidate");
  if (e.candidate) {
    console.log(e.candidate);
    websocketRef.value.send(JSON.stringify({ iceCandidate: e.candidate }));
  }
};

const handleTrackEvent = (e) => {
  console.log("received track");
  partnerVideo.value.srcObject = e.streams[0];
};

const handleOffer = async (offer) => {
  console.log("received offer, creating answer");
  
  peerRef.value = createPeer();
  await peerRef.value.setRemoteDescription(new RTCSessionDescription(offer));
  
  userStream.value.getTracks().forEach((track) => {
    peerRef.value.addTrack(track, userStream.value);
  });
  
  const answer = await peerRef.value.createAnswer();
  await peerRef.value.setLocalDescription(answer);
  
  websocketRef.value.send(JSON.stringify({ answer: peerRef.value.localDescription }));
};

onMounted(() => {
  openCamera().then((stream) => {
    if (stream) {
      userVideo.value.srcObject = stream;
      userStream.value = stream;
      
      websocketRef.value = new WebSocket(BASE_URL);
      
      websocketRef.value.addEventListener("open", () => {
        websocketRef.value.send(JSON.stringify({ join: true }));
      });
      
      websocketRef.value.addEventListener("message", async (e) => {
        const message = JSON.parse(e.data);
        
        if (message.join) {
          callUser();
        }
        
        if (message.offer) {
          handleOffer(message.offer);
        }
        
        if (message.answer) {
          console.log("receiving answer");
          await peerRef.value.setRemoteDescription(new RTCSessionDescription(message.answer));
        }
        
        if (message.iceCandidate) {
          console.log("receiving and adding ice candidate");
          try {
            await peerRef.value.addIceCandidate(message.iceCandidate);
          } catch (err) {
            console.log("Error receiving ice candidate:", err);
          }
        }
      });
    } else {
      console.error("Failed to open camera.");
    }
  });
});
</script>