import { createRouter, createWebHistory } from 'vue-router';
import Room from '../components/Room.vue';
import CreateRoom from '../components/CreateRoom.vue'

const routes = [
  { 
    path: '/', 
    name: 'CreateRoom',
    component: CreateRoom 
  },
  { 
    path: '/room/:roomID', 
    name: 'Room',
    component: Room 
  },

];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
