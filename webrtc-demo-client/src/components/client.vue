<template>
<div class="wrapper">
    <div class="top">
        <h2 style="margin-right: 20px;">呼叫方</h2>
        <el-button @click="getUserMedia" type="primary">1. 点击共享</el-button>
        <el-button @click="call" type="primary">2.呼叫</el-button>
    </div>
    <el-row :gutter="10">
        <el-col :xs="6" :md="6" :lg="6" :xl="6">
            <el-card style="height: 30vh;">
                <template #header>当前共享的区域</template>
                <video ref="localVideo" autoplay width="100%" height="100%" :srcObject="localStream"></video>
            </el-card>
        </el-col>
        <el-col :xs="18" :md="18" :lg="18" :xl="18">
            <el-card style="height: 80vh;">
                <template #header>远程视频区域</template>
                <video ref="remoteVideo" autoplay width="100%" height="100%" :srcObject="remoteStream"></video>
            </el-card>
        </el-col>
      </el-row>
</div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
const baseURL = import.meta.env.VITE_WS_BASEURL
const peer = new RTCPeerConnection()
const localVideo = ref()
const remoteVideo = ref()
const localStream = ref()
const remoteStream = ref()
const roomId = ref(1)
const remoteSdp = ref<{type:'', sdp: ''}>({type: '', sdp: ''})
const candidate = ref<RTCIceCandidate>()
const route = useRoute()
const userId = route.params.id 
const ws = new WebSocket(`ws://${baseURL}/rtc?roomId=${roomId.value}&userId=${userId}`)
if(ws) {
    console.log("已连接")
    ws.addEventListener('message', (e)=>{
        const data =JSON.parse(e.data)
        if(data.key === 'answer_sdp') {
            remoteSdp.value = data.sdp
            setRemoteDescription()
        }
        if(data.key === 'answer_candidate') {
            candidate.value = data.candidate
            setAnswerCandidate()
        }
    })
}

const getUserMedia = async()=>{
    const stream = await navigator.mediaDevices.getDisplayMedia({audio: true, video: true
    })
    stream.getTracks().forEach((track)=>{
        peer.addTrack(track, stream)
    })
    localStream.value = stream

    const offer = await peer.createOffer()
    peer.setLocalDescription(offer)
    const params = {
        key: "offer_sdp",
        sdp: offer
    }
    ws.send(JSON.stringify(params))
}

const call = ()=>{
    const params = {
        key: 'call_remote',
        data: {}
    }
    ws.send(JSON.stringify(params))
}

peer.ontrack = e => {
      remoteStream.value = e.streams[0];
      remoteVideo.value.oncanplay = () => {
        remoteVideo.value.play();
    };
};
peer.addEventListener('icecandidate', (e)=>{
    if(e.candidate) {
        candidate.value = JSON.stringify(e.candidate) as any
        const params = {
            key: 'offer_candidate',
            candidate: e.candidate
        }
        ws.send(JSON.stringify(params))
    }
})

const createOffer = ()=>{
}

const setRemoteDescription = () => {
    const des = {type: 'answer', sdp: remoteSdp.value.sdp }
    peer.setRemoteDescription(des as any) 

}

const setAnswerCandidate = () => {
    peer.addIceCandidate(candidate.value)
}

</script>

<style lang="css" scoped>
.wrapper {
    width: 90vw;
    padding: 0 16px;
}
.top {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
}
.video-wrapper {
    display: flex;
}
.video-container {
    border: 1px solid gray;
}
.title {
    height: 20%;
    width: 100%;
}
</style>