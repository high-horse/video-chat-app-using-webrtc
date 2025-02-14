
<template>
    <div>
        <video autoplay controls="true" ref="userVideo"></video>
        <video autoplay controls="true" ref="clientVideo"></video>

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
const clientVideo = ref(null);
const peerRef = ref(null);
const websocketRef = ref(null);

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

onMounted(() => {
  openCamera().then((stream) => {
    // userVideo.value.current.srcObject = stream;
    userVideo.value.srcObject = stream;
    
    userStream.value = stream;
  });
})

</script>
