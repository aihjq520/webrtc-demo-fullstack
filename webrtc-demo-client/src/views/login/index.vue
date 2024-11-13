<template>
  <div class="login-container">
    <el-card class="login-card">
      <div slot="header" class="login-header">
        <span>登录</span>
      </div>
      <el-form :model="loginForm" ref="loginFormRef" class="login-form" :rules="rules" label-position="left" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="loginForm.password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-radio-group v-model="loginForm.type">
            <el-radio value="call" size="large">呼叫方</el-radio>
            <el-radio value="accept" size="large">接收方</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { reactive, ref, unref } from 'vue'
import router from '../../routers';

const loginForm = reactive({
        username: '',
        password: '',
        type: 'call'
      })
      
      const rules = {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ],
        type: [
          { required: true }
        ]
      }
  
      const loginFormRef = ref(null)
      const baseURL = import.meta.env.VITE_WS_BASEURL
      const handleLogin = () => {
        loginFormRef.value.validate(async(valid) => {
          if (valid) {
            const params = {
              telephone: loginForm.username,
              password: loginForm.password
            }
            // 在这里处理登录逻辑，例如发送请求到服务器
            const response = await fetch(`http://${baseURL}/api/auth/login`, {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify(params)
            });
            if(response.status === 200) {
              const res = await response.json()
              router.replace(loginForm.type === 'call' ? '/offer' : '/answer', {params: {id: res.data.id}})
            } else{
              alert('登录失败')
            }
          } else {
            console.log('error submit!!')
            return false
          }
        })
      }
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: radial-gradient(
    circle at 60% 90%,
    rgba(46, 103, 161, 1),
    transparent 60%
    ),
    radial-gradient(
    circle at 20px 20px,
    rgba(46, 103, 161, 0.8),
    transparent 25%
    ),
    #182336
}

.login-card {
  width: 400px;
}

.login-header {
  text-align: center;
}

.login-form {
  margin-top: 20px;
}
</style>
