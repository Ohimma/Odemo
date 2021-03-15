<template>
  <div class="login">
    <div class="login-msg">
      <div class="title">
        <img src="../assets/img/login-logo.png" alt="">
        <span>管理后台</span>
      </div>
      <!-- 
        :model指定表单使用的数据 :rules进行表单校验规则 ref是提交的对象名字
        prop表单校验的属性
        input 绑定的元素必须是model的直接属性 
        prefix 左边 suffix 右边 clearable 删除所有输入 
        :loading  true时会转圈加载 false时正常
        status-icon 添加校验结果反馈显示图标
      -->
      <el-form :model="loginData" :rules="loginDataRules" ref="loginForm">
        <el-form-item prop="account">
          <el-input type="text" 
            v-model="loginData.account" 
            prefix-icon="el-icon-user" 
            placeholder="请输入用户名"
            tabindex="1"
            auto-complete="on">
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginData.password"
            type="password"
            placeholder="请输入密码"  
            prefix-icon="el-icon-unlock"
            clearable show-password
            
            @keyup.enter="loginBtn"
            auto-complete="on">
          </el-input>
        </el-form-item>

        <el-form-item>
          <el-button 
            type="primary"
            :loading="loading" 
            @click.prevent="loginBtn">登&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      loginData: {
        account: '',
        password: ''
      },
      loginDataRules: {
        account: {required: true, message: '请输入用户名', trigger: 'change'},
        password: [{required: true, message: '请输入密码', trigger: 'change'}]
      },
      loading: false,
    }
  },
  created(){
    console.log("entrey login create")
  },
  mounted(){
    console.log("entrey login mounted")
    this.loginData.account = localStorage.getItem('userName')
  },
  methods: {
    loginBtn () {
      this.loading = true
      // 后端要求：先拿一个rsa公钥，对传输的密码进行加密传输，公钥是一次性的
      this.$http.get("http://localhost:8080/api/console/auth/key")
        .then((res) => {
          this.handleLogin(res.data.key)
        })
        .catch((err) => {
          this.$message({type: 'error', message: "获取加密公钥失败"})
          this.loading = false
          console.log('get errr', err)
        })
    },
    handleLogin: function (authKey) {
      // validate 对整个表单验证，参数为回调函数，不传入函数返回promise
      this.$refs.loginForm.validate(valid => {
        if (valid) {  // vaild 是 true/false
          // post 传入用户名和 rsa加密后的用户密码
          this.$http.post('http://localhost:8080/api/console/auth/login', {
            account: this.loginData.account,
            password: this.$utils.encrypt(this.loginData.password, authKey)
          }, false)
            .then((res) => {
              this.$message({
                type: 'success',
                message: '登陆成功'
              })
              this.$store.dispatch('USER_TOKEN_ACTION', res.data.token)
              this.$store.dispatch('USER_NAME_ACTION', this.loginData.account)
              // 登陆成功后，判断是否有redirect，没有的话跳转到/home
              // decodeURIComponent对encodeURLComponent的url乱码解码，
              let redirect = decodeURIComponent(this.$route.query.redirect || '/home')
              this.$router.push({
                path: redirect
              })
              this.loading = false
            })
            .catch((err) => {
              this.loading = false
              this.$message({
                type: 'error', 
                message: "认证失败"
              })
              
            })
        } else {
          this.loading = false
        }
      })
    }
  }
}
</script>

<style scoped>
.login {
  display: flex;
  justify-content: center;
  align-items: center;

  height: 100%;
  width: 100%;

  background-size: 100%;
  overflow: hidden;
  background: #11255A url('../assets/img/login-bg.png') no-repeat center bottom;

  transition: all 1s;
}
.login .login-msg {
    width: 400px;
    height: 400px;
    padding: 20px;
    box-sizing: border-box;
    /* background-color: pink; */
}
.login .login-msg .title {
    width: 100%;
    height: 25%;

    display: flex;
    justify-content: center;
    align-items: center;

    margin-bottom: 20px;
}
.login .login-msg .title img {
    width: 40px;
    height: 40px;
}
.login .login-msg .title span {
  font-size: 20px;
  line-height: 50px;
  color: #FFFFFF;
  margin: 0px 10px;
}

.el-button {
  width: 100%;
  height: 40px;
  margin-top: 25px;
  
  background: #204CC0;
  border: 0;
  border-radius: 4px;

  color: #fff;
  font-size: 20px;
}
.el-button:hover,
.el-button.is-fouse {
  background: #4d73da;
} 


</style>

<style >
.login .el-form .el-input__inner {
  padding-left: 40px;
  letter-spacing: 2px;
  font-size: 16px;
  font-family: "MingLiu";
}
</style>