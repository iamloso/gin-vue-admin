<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="通讯地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市:">
          <el-input v-model="formData.city" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="申报条件:">
          <el-input v-model="formData.conditions" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所在单位:">
          <el-input v-model="formData.currentUnit" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出生日期:">
          <el-input v-model="formData.dateBirth" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文化程度:">
          <el-input v-model="formData.eduLevel" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input v-model="formData.email" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简要经历:">
          <el-input v-model="formData.experience" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证件类型:">
          <el-input v-model="formData.IDType" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="级别:">
          <el-input v-model="formData.level" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="民族:">
          <el-input v-model="formData.nation" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="原证书等级:">
          <el-input v-model="formData.OCL" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="原证书编号:">
          <el-input v-model="formData.OCN" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="原证书职业:">
          <el-input v-model="formData.OCO" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="户籍所在地:">
          <el-input v-model="formData.place" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="政治面貌:">
          <el-input v-model="formData.politicalStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮寄地址:">
          <el-input v-model="formData.postAddress" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮政编码:">
          <el-input v-model="formData.postalCode" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="职业名称:">
          <el-input v-model="formData.professionalName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="省份:">
          <el-input v-model="formData.province" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证件领取:">
          <el-input v-model="formData.receive" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学历证书编号:">
          <el-input v-model="formData.serialNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-input v-model="formData.sex" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="考生来源:">
          <el-input v-model="formData.source" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号:">
          <el-input v-model="formData.UID" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="从事专业:">
          <el-input v-model="formData.work" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="参加工作时间:">
          <el-input v-model="formData.workDate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="工种名称:">
          <el-input v-model="formData.workType" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专业年限:">
          <el-input v-model="formData.workYear" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'JyxUser'
}
</script>

<script setup>
import {
  createJyxUser,
  updateJyxUser,
  findJyxUser
} from '@/api/jyxUser'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        address: '',
        city: '',
        conditions: '',
        currentUnit: '',
        dateBirth: '',
        eduLevel: '',
        email: '',
        experience: '',
        IDType: '',
        level: '',
        name: '',
        nation: '',
        OCL: '',
        OCN: '',
        OCO: '',
        phone: '',
        place: '',
        politicalStatus: '',
        postAddress: '',
        postalCode: '',
        professionalName: '',
        province: '',
        receive: '',
        serialNumber: '',
        sex: '',
        source: '',
        UID: '',
        work: '',
        workDate: '',
        workType: '',
        workYear: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findJyxUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rejyxUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createJyxUser(formData.value)
          break
        case 'update':
          res = await updateJyxUser(formData.value)
          break
        default:
          res = await createJyxUser(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
