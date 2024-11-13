<template>
    <div class="wrapper">
    <div class="top">
        <el-button @click="createAnswer" type="primary">点击共享</el-button>
        <h2 style="margin-left: 20px;">接受方</h2>
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
    import { ElMessage, ElMessageBox } from 'element-plus'
    const peer = new RTCPeerConnection()
    
    
    const localVideo = ref()
    const localStream = ref()
    const remoteStream = ref()
    const remoteVideo = ref()
    const roomId = ref(1)
    const remoteSdp = ref<{type: string, sdp: string}>({type: '', sdp: ''})
    const candidate = ref<RTCIceCandidate>()
    const getUserMedia = async()=>{
       
        const stream = await navigator.mediaDevices.getDisplayMedia({audio: true, video: true
        })
        stream.getTracks().forEach((track)=>{
            peer.addTrack(track, stream)
        })

        return stream
        // answer candidate
    }
    const baseURL = import.meta.env.VITE_WS_BASEURL
    const ws = new WebSocket(`ws://${baseURL}/rtc?roomId=${roomId.value}&userId=3`) // 房间号
    if(ws) {
        console.log("已连接")
        ws.addEventListener('message', (e)=>{
            console.log(e.data)
            const data =JSON.parse(e.data)
            if(data.key==='call_remote') {
                // 弹出提醒
                acceptCall(data)
            }
            if(data.key === 'offer_sdp') {
                remoteSdp.value = data.sdp
                setRemoteDescription()
            }
            if(data.key === 'offer_candidate') {
                candidate.value = data.candidate
                setOfferCandidate()
            }
        })
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
            key: 'answer_candidate',
            candidate: e.candidate
        }
        ws.send(JSON.stringify(params))
    }
    })
    
    const createAnswer = async ()=>{
        
        localStream.value = await getUserMedia()
        setRemoteDescription()
        
        const answer = await peer.createAnswer()
        peer.setLocalDescription(answer)
        const params = {
            key: "answer_sdp",
            sdp: answer
        }
        ws.send(JSON.stringify(params))
    }
    
    const setRemoteDescription = () => {
        const des = { type: 'offer', sdp: remoteSdp.value.sdp }
        peer.setRemoteDescription(des as any) 
    
    }
    
    const setOfferCandidate = () => {
        peer.addIceCandidate(candidate.value)
    }

    const acceptCall = (data) => {
        ElMessageBox.confirm(
            '你的好朋友发来呼叫消息',
            'Warning',
            {
            confirmButtonText: '接听',
            cancelButtonText: '拒绝',
            type: 'warning',
            }
        )
        .then(() => {
            createAnswer()
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '已拒绝',
            })
        })
    }
    
    </script>

<style scoped>
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