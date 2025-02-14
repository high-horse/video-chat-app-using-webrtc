<script setup>
import { QBtn } from 'quasar';  // Import QBtn explicitly
import {useRouter} from 'vue-router';
import api from '../utils/api.js'

const router = useRouter()
const handleCreateRoom = async () => {
  const response = await api.get('create');
  console.log(response.data)
  const roomId = response.data.room_id
  router.push(`room/${roomId}`)
}
</script>

<template>
  <div>
    <q-btn label="Create Room" color="primary" class="capitalize" @click.prevent=handleCreateRoom />
  </div>
</template>
