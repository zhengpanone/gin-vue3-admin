<template>
  <div class="login-container">
    <el-form
        class="login-form"
        :rules="rules"
        ref="form"
        :model="user"
        @submit.prevent="handleSubmit"
    >
      <div class="login-form-header">
        <img class="login-logo" src="@/assets/login_logo.png" alt="logo"/>
      </div>
      <el-form-item prop="username">
        <el-input v-model="user.username" placeholder="请输入用户名">
          <template #prefix>
            <i class="el-input_icon">
              <el-icon>
                <User/>
              </el-icon>
            </i>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
            v-model="user.password"
            type="password"
            placeholder="请输入密码"
        >
          <template #prefix>
            <i class="el-input__icon">
              <el-icon>
                <Lock/>
              </el-icon>
            </i>
          </template>
        </el-input>
      </el-form-item>

      <el-form-item prop="captcha">
        <div class="imgcode-wrap">
          <el-input v-model="user.captcha" placeholder="请输入验证码">
            <template #prefix>
              <i class="el-input_icon">
                <el-icon>
                  <Key/>
                </el-icon>
              </i>
            </template>
          </el-input>
          <img
              class="imgcode"
              alt="验证码"
              :src="captchaSrc"
              @click="loadCaptcha"
          />
        </div>
      </el-form-item>
      <el-form-item>
        <el-button
            class="submit-button"
            type="primary"
            :loading="loading"
            native-type="submit"
        >
          登录
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script setup lang="ts">
import {Key, Lock, User} from "@element-plus/icons-vue";
import {getCaptcha, login} from "@/api/common";
import type {ICaptchaInfo} from "@/api/types/common";
import {onMounted, reactive, ref, toRaw} from "vue";
import type {IElForm, IFormRule} from "@/types/element-plus";
import {useRouter} from "vue-router";
import {indexStore} from "@/store/index";
import {setItem} from "@/utils/storage";

const router = useRouter();
const form = ref<IElForm | null>(null);

const captchaSrc = ref<ICaptchaInfo["pic_path"]>();
const user = reactive({
  username: "admin",
  password: "admin",
  captcha: "",
  captchaId: "",
});
onMounted(() => {
  loadCaptcha();
});

const loadCaptcha = async () => {
  await getCaptcha().then((res) => {
    captchaSrc.value = res.data.pic_path;
    user.captchaId = res.data.captcha_id;
  });
};

const loading = ref(false);
const rules = ref<IFormRule>({
  username: [{required: true, message: "请输入账号", trigger: "change"}],
  password: [{required: true, message: "请输入密码", trigger: "change"}],
  captcha: [{required: true, message: "请输入验证码", trigger: "change"}],
});

const handleSubmit = async () => {
  // 表单验证
  const valid = await form.value?.validate();
  if (!valid) {
    return false;
  }
  // 验证通过,展示loading
  loading.value = true;
  // 请求提交

  try {
    const loginData = await login(toRaw(user));
    const store = indexStore();

    store.setUser({...loginData.data.userInfo, token: loginData.data.token});
    await router.replace({
      name: "home",
    });
    // 处理响应
    console.log("handleSubmit");
  } catch (error) {
    loadCaptcha();
  } finally {
    loading.value = false;
  }
};
</script>

<style lang="scss" scoped>
.login-container {
  min-width: 400px;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #2d3a4b;
}

.login-form {
  padding: 30px;
  border-radius: 6px;
  background: #fff;
  min-width: 350px;

  .login-form__header {
    display: flex;
    justify-content: center;
    align-items: center;
    padding-bottom: 30px;
  }

  .el-form-item:last-child {
    margin-bottom: 0;
  }

  .login__form-title {
    display: flex;
    justify-content: center;
    color: #fff;
  }

  .submit-button {
    width: 100%;
  }

  .login-logo {
    width: 271px;
    height: 74px;
  }

  .imgcode-wrap {
    display: flex;
    align-items: center;
    width: 100%;

    .imgcode {
      height: 37px;
    }
  }
}
</style>
