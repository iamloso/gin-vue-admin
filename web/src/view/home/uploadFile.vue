<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <img style="float: left" src="./logo.png">
        <span style="float: left; font-size: 25px">师范教育系在线报名</span>
      </el-header>
      <el-main>
        <el-form ref="elForm" :model="formData" :rules="rules" size="small" label-width="100px">
          <el-scrollbar height="600px">
            <el-form-item label="姓名" prop="name">
              <el-input
                v-model="formData.name"
                placeholder="请输入姓名"
                clearable
                :style="{width: '100%'}"
              />
            </el-form-item>
            <el-form-item size="large">
              <el-button type="primary" @click="submitForm">提交</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-scrollbar>
        </el-form>
      </el-main>
      <el-footer>版权所有© Copyright 2022-2025 北京诗码作集科技有线公司技术支持(电话：18513172626)</el-footer>
    </el-container>
  </div>
</template>

<script>
import { createJyxUser } from '@/api/jyxUser'
import { ElMessage } from 'element-plus'

export default {
  name: 'UploadFile',
  components: {},
  props: [],
  data() {
    return {
      formData: {
        name: undefined,
      },
      rules: {
        name: [{
          required: true,
          message: '请输入姓名',
          trigger: 'blur'
        }],
      },
    }
  },
  computed: {},
  watch: {},
  created() { console.log(123) },
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(async valid => {
        if (!valid) return
        // TODO 提交表单
        const res = await createJyxUser(JSON.stringify(this.formData))
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '创建/更改成功'
          })
        }
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
  }
}
</script>

<style scoped>
/*页面布局*/
.common-layout .el-header,
.common-layout .el-footer {
  background-color: #409EFF;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
}
.common-layout .el-footer {
  line-height: 60px;
}

.common-layout .el-aside {
  background-color: #d3dce6;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 200px;
}

.common-layout .el-main {
  background-color: #e9eef3;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 160px;
}

.common-layout > .el-container {
  margin-bottom: 40px;
}

.common-layout .el-container:nth-child(5) .el-aside,
.common-layout .el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.common-layout .el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
/*页面布局*/

/*滚动条*/
.scrollbar-demo-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  margin: 10px;
  text-align: center;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
}
/*滚动条*/
</style>
