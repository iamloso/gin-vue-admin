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
            <el-form-item label="报名项目" prop="professionalName">
              <span style="float:left;">{{formData.professionalName}}</span>
            </el-form-item>
            <el-form-item label="姓名" prop="name">
              <span style="float:left;">{{formData.name}}</span>
            </el-form-item>
            <el-form-item label="手机号" prop="phone">
              <span style="float:left;">{{formData.phone}}</span>
            </el-form-item>
            <el-form-item label="证件号" prop="UID">
              <span style="float:left;">{{formData.UID}}</span>
            </el-form-item>
            <el-form-item label="照片" prop="name">
              <img style="float: left" :src="`${path}`+'/'+formData.userPic" width="80">
            </el-form-item>
            <el-form-item label="扫码支付" prop="isPay">
              <span style="float:left;">{{formData.isPay}}</span>
            </el-form-item>
            <el-form ref="elForm" :model="formData" :rules="rules" size="small" label-width="100px">
              <el-form-item label="支付金额" prop="payAmount">
                <span style="float:left;">{{formData.payAmount}}</span>
              </el-form-item>
            </el-form>
            <el-form-item label="支付凭证" prop="name">
              <img style="float: left" :src="`${path}`+'/'+formData.userPay" width="80">
            </el-form-item>
          </el-scrollbar>
        </el-form>
      </el-main>
      <el-footer></el-footer>
    </el-container>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const path = ref(import.meta.env.VITE_BASE_API)
</script>

<script>
import { getJyxUserList } from '@/api/jyxUser'

export default {
  name: 'Home',
  components: {},
  props: [],
  data() {
    return {
      formData: {
        // CreatedAt: undefined,
        // UpdatedAt: undefined,
        ID: undefined,
        name: undefined,
        IDType: undefined,
        UID: undefined,
        sex: undefined,
        dateBirth: null,
        eduLevel: undefined,
        source: undefined,
        receive: undefined,
        professionalName: undefined,
        workType: undefined,
        level: undefined,
        conditions: undefined,
        phone: undefined,
        workDate: undefined,
        work: undefined,
        workYear: undefined,
        nation: undefined,
        politicalStatus: undefined,
        serialNumber: undefined,
        experience: undefined,
        postalCode: undefined,
        currentUnit: undefined,
        address: undefined,
        postAddress: undefined,
        email: undefined,
        place: undefined,
        province: undefined,
        city: undefined,
        OCO: undefined,
        OCL: undefined,
        OCN: undefined,
        userPic: undefined,
        userCertify: undefined,
        userPay: undefined,
        isPay: undefined,
        payAmount: undefined,
      },
      rules: {
        name: [{
          required: true,
          message: '请输入姓名',
          trigger: 'blur'
        }],
        IDType: [{
          required: true,
          message: '请输入证件类型',
          trigger: 'change'
        }],
        UID: [{
          required: true,
          message: '请输入证件号',
          trigger: 'blur'
        }],
        sex: [{
          required: true,
          message: '请选择性别',
          trigger: 'change'
        }],
        dateBirth: [{
          required: true,
          message: '请选择出生日期',
          trigger: 'change'
        }],
        eduLevel: [{
          required: true,
          message: '请选择文化程度',
          trigger: 'change'
        }],
        source: [{
          required: true,
          message: '请选择考生来源',
          trigger: 'change'
        }],
        receive: [{
          required: true,
          message: '请选择证书领取',
          trigger: 'change'
        }],
        work: [{
          required: true,
          message: '请选择职业',
          trigger: 'change'
        }],
        professionalName: [{
          required: true,
          message: '请选择职业名称',
          trigger: 'change'
        }],
        workType: [],
        level: [{
          required: true,
          message: '请选择年级',
          trigger: 'change'
        }],
        conditions: [{
          required: true,
          message: '请选择班级',
          trigger: 'change'
        }],
        phone: [{
          required: true,
          message: '请输入手机号',
          trigger: 'blur'
        }],
        workDate: [],
        workYear: [],
        nation: [{
          required: true,
          message: '请选择民族',
          trigger: 'change'
        }],
        politicalStatus: [],
        serialNumber: [],
        experience: [],
        postalCode: [],
        currentUnit: [{
          required: true,
          message: '请选择所在单位',
          trigger: 'change'
        }],
        address: [],
        postAddress: [],
        email: [],
        place: [],
        province: [{
          required: true,
          message: '请选择省份',
          trigger: 'change'
        }],
        city: [{
          required: true,
          message: '请选择城市',
          trigger: 'change'
        }],
        OCO: [],
        OCL: [],
        OCN: [],
        isPay: [{
          required: true,
          message: '扫码支付不能为空',
          trigger: 'change'
        }],
        payAmount: [{
          required: true,
          message: '请输入金额支付支付金额',
          trigger: 'blur'
        }],
      },
    }
  },
  computed: {},
  watch: {},
  async created() {
    if (window.localStorage.getItem('UID')) {
      this.formData.ID = window.localStorage.getItem('UID')
      const res = await getJyxUserList(JSON.stringify(this.formData))
      if (res.code === 0 && res.data.list.length > 0) {
        console.log(res.data.list[0])
        this.formData = res.data.list[0]
        // await this.$router.push({ name: 'UploadFile', params: { name: this.formData.name, UID: this.formData.UID }})
      }
    }
  },
  mounted() {},
  methods: {}
}
</script>

<style scoped>
.upload-btn+.upload-btn {
  margin-left: 12px;
}

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
