
<template>
    <div>
        <video autoplay controls="true" ref="userVideo"></video>
    <video autoplay controls="true" ref="partnerVideo">m</video>

        <!-- <q-video autoplay control='true'></q-video> -->
    </div>
</template>


<script setup>
// import QVideo from 'quasar';
import { ref } from "vue";
import { onMounted } from "vue";
import { useRoute } from "vue-router";

const router = useRoute();

const userVideo = ref(null);
const userStream = ref(null);
const partnerVideo = ref(null);
const peerRef = ref(null);
const websocketRef = ref(null);

const BASE_URL = `ws://localhost:8000/join?roomID=${router.params.roomID}`

const openCamera = async() => {
  const allDevices = await navigator.mediaDevices.enumerateDevices();
  const cameras = allDevices.filter(
    (device) => device.kind == "videoinput"
  );
  console.log('cameras', cameras)
  const contrained = {
    audio: true,
    video: {
      // cursor: "always"
      // deviceId: cameras[1].deviceId,
    }
  }
  
  try{
    return await navigator.mediaDevices.getUserMedia(contrained);
  }catch(err) {
    console.error(err);
  }
}

const callUser = () => {
  console.log("calling other user");
  peerRef.value =createPeer();
  
  userStream.value.getTracks().forEach((track) => {
    peerRef.value.addTrack(track, userStream.value )
  })
}
const createPeer = () => {
  console.log("creating peer connection");
  
  const peer = new RTCPeerConnection({
    iceServers: [{ urls: "stun:stun.l.google.con:19302" }]
  })
  
  peer.onnegotiationneeded = handleNegotiationNeeded;
  peer.onicecandidate = handleIceCandidateEvent();
  peer.ontrack = handleTrackEvent();
  
  return peer;
}

const handleNegotiationNeeded = async() => {
  console.log("creating offer");
  try {
    const myOffer = await peerRef.value.createOffer();
    await peerRef.value.setLocalDescription(myOffer);
    
    // websocketRef.value.send(JSON.stringify({ offer: myOffer }));
    websocketRef.value.send(JSON.stringify({ offer: peerRef.value.localDescription }));
  }catch(err) {
    
  }
}
const handleIceCandidateEvent = (e) => {
  console.log("found ice candidate");
  if(e.candidate) {
    console.log(e.candidate);
    websocketRef.value.send(JSON.stringify({iceCandidate: e.candidate}));
  }
}

const handleTrackEvent = (e) => { 
  console.log("Recieve track");
  partnerVideo.value.srcObject = e.streams[0];
}

const handleOffer = async(offer) => {
  console.log("recieved offer creating answer")
  
  peerRef.value = createPeer();
  // const desv = new RTCSessionDescription(offer)
  await peerRef.value.setRemoteDescription(new RTCSessionDescription(offer));
  
  userStream.value.getTracks().forEach((track) => {
    peerRef.value.addTrack(track, userStream.value);
  });
  
  const answer = await peerRef.value.createAnswer();
  await peerRef.value.setLocalDescription(answer);
  
  websocketRef.value.send(JSON.stringify({ answer: peerRef.value.localDescription }));
  
}

onMounted(() => {
  openCamera().then((stream) => {
    userVideo.value.srcObject = stream;
    userStream.value = stream;
    
    websocketRef.value = new WebSocket(BASE_URL);
    
    websocketRef.value.addEventListener("open", () => {
      websocketRef.value.send(JSON.stringify({join: true}));
    })
    
    websocketRef.value.addEventListener("message" ,async (e) => {
      const message = JSON.parse(e.data);
      
      if(message.join) {
        callUser();
      }
      
      if(message.offer) {
        handleOffer(message.offer)
      }
      
      if(message.answer) {
        console.log("recieving  answer");
        peerRef.value.setRemoteDescription(new RTCSessionDescription(message.answer));
      }
      
      if(message.iceCandidate) {
        console.log("recieving and adding ice candidate");
        try{
          await peerRef.value.addIceCandidate(message.iceCandidate)
        }catch(err) {
          console.log("Error recieving ice candidate :", err)
        }
      }
      
    })
  });
})

</script>
