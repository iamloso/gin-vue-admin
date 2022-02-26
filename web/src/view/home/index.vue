<template>
  <div>
    <el-form ref="elForm" :model="formData" :rules="rules" size="small" label-width="100px">
      <el-form-item label="姓名" prop="name">
        <el-input
          v-model="formData.name"
          placeholder="请输入姓名"
          :maxlength="11"
          show-word-limit
          clearable
          prefix-icon="el-icon-mobile"
          :style="{width: '100%'}"
        />
      </el-form-item>
      <el-form-item label="证件类型" prop="IDType">
        <el-select v-model="formData.IDType" placeholder="请输入证件类型" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in IDTypeOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="证件号" prop="ID">
        <el-input v-model="formData.ID" placeholder="请输入证件号" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="性别" prop="sex">
        <el-select v-model="formData.sex" placeholder="请选择性别" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in sexOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="出生日期" prop="dateBirth">
        <el-date-picker
          v-model="formData.dateBirth"
          format="yyyy-MM-dd"
          value-format="yyyy-MM-dd"
          :style="{width: '100%'}"
          placeholder="请选择出生日期"
          clearable
        />
      </el-form-item>
      <el-form-item label="文化程度" prop="eduLevel">
        <el-select v-model="formData.eduLevel" placeholder="请选择文化程度" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in eduLevelOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="考生来源" prop="field119">
        <el-select v-model="formData.field119" placeholder="请选择考生来源" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in field119Options"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="证书领取" prop="field120">
        <el-select v-model="formData.field120" placeholder="请选择证书领取" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in field120Options"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="职业名称" prop="professionalName">
        <el-select
          v-model="formData.professionalName"
          placeholder="请选择职业名称"
          clearable
          :style="{width: '100%'}"
        >
          <el-option
            v-for="(item, index) in professionalNameOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="工种名称" prop="workType">
        <el-select v-model="formData.workType" placeholder="请选择工种名称" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="级别" prop="level">
        <el-select v-model="formData.level" placeholder="请选择级别" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in levelOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="申请条件" prop=" conditio">
        <el-select v-model="formData. conditio" placeholder="请选择申请条件" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in conditioOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="formData.phone" placeholder="请输入手机号" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="参加工作时间" prop="workDate">
        <el-input v-model="formData.workDate" placeholder="请输入参加工作时间" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="从事专业" prop="work">
        <el-input v-model="formData.work" placeholder="请输入从事专业" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="专业年限" prop="workYear">
        <el-input v-model="formData.workYear" placeholder="请输入专业年限" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="民族" prop="nation">
        <el-select v-model="formData.nation" placeholder="请选择民族" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in nationOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="政治面貌" prop="politicalStatus">
        <el-select
          v-model="formData.politicalStatus"
          placeholder="请选择政治面貌"
          clearable
          :style="{width: '100%'}"
        >
          <el-option
            v-for="(item, index) in politicalStatusOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="学历证书编号" prop="serialNumber">
        <el-input v-model="formData.serialNumber" placeholder="请输入学历证书编号" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="简要经历" prop="experience">
        <el-input
          v-model="formData.experience"
          type="textarea"
          placeholder="请输入简要经历"
          :autosize="{minRows: 4, maxRows: 4}"
          :style="{width: '100%'}"
        />
      </el-form-item>
      <el-form-item label="邮政编号" prop="postalCode">
        <el-input v-model="formData.postalCode" placeholder="请输入邮政编号" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="所在单位" prop="currentUnit">
        <el-select v-model="formData.currentUnit" placeholder="请选择所在单位" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in currentUnitOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="通讯地址" prop="address">
        <el-input v-model="formData.address" placeholder="请输入通讯地址" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="邮寄地址" prop="postAddress">
        <el-input v-model="formData.postAddress" placeholder="请输入邮寄地址" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="formData.email" placeholder="请输入邮箱" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="户籍所在地" prop="place">
        <el-input v-model="formData.place" placeholder="请输入户籍所在地" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="省份" prop="province">
        <el-select v-model="formData.province" placeholder="请选择省份" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in provinceOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="城市" prop="city">
        <el-select v-model="formData.city" placeholder="请选择城市" clearable :style="{width: '100%'}">
          <el-option
            v-for="(item, index) in cityOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="原证书职业" prop="OCO">
        <el-input v-model="formData.OCO" placeholder="请输入原证书职业" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="原证书等级" prop="OCL">
        <el-input v-model="formData.OCL" placeholder="请输入原证书等级" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="原证书编号" prop="OCN">
        <el-input v-model="formData.OCN" placeholder="请输入原证书编号" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  components: {},
  props: [],
  data() {
    return {
      formData: {
        name: undefined,
        IDType: 1,
        ID: undefined,
        sex: undefined,
        dateBirth: null,
        eduLevel: undefined,
        field119: undefined,
        field120: undefined,
        professionalName: undefined,
        workType: undefined,
        level: undefined,
        conditio: undefined,
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
      },
      rules: {
        name: [{
          required: true,
          message: '请输入姓名',
          trigger: 'blur'
        }, {
          pattern: /^1(3|4|5|7|8|9)\d{9}$/,
          message: '手机号格式错误',
          trigger: 'blur'
        }],
        IDType: [{
          required: true,
          message: '请输入证件类型',
          trigger: 'change'
        }],
        ID: [{
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
        field119: [{
          required: true,
          message: '请选择考生来源',
          trigger: 'change'
        }],
        field120: [{
          required: true,
          message: '请选择证书领取',
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
          message: '请选择级别',
          trigger: 'change'
        }],
        conditio: [{
          required: true,
          message: '请选择申请条件',
          trigger: 'change'
        }],
        phone: [{
          required: true,
          message: '请输入手机号',
          trigger: 'blur'
        }],
        workDate: [],
        work: [],
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
      },
      IDTypeOptions: [{
        'label': '身份证',
        'value': 1
      }],
      sexOptions: [{
        'label': '男',
        'value': 1
      }, {
        'label': '女',
        'value': 2
      }],
      eduLevelOptions: [{
        'label': '大学专科和专科学校',
        'value': 1
      }, {
        'label': '大学本科',
        'value': 2
      }, {
        'label': '研究生',
        'value': 3
      }],
      field119Options: [{
        'label': '机关事业单位',
        'value': 1
      }, {
        'label': '职业技术学院',
        'value': 2
      }],
      field120Options: [{
        'label': '自取',
        'value': 1
      }],
      professionalNameOptions: [{
        'label': '保育员',
        'value': 1
      }],
      levelOptions: [{
        'label': '三级',
        'value': 1
      }, {
        'label': '高级工',
        'value': ''
      }],
      conditioOptions: [{
        'label': '2',
        'value': 1
      }],
      nationOptions: [{
        'label': '蒙古族',
        'value': 1
      }, {
        'label': '汉族',
        'value': 2
      }, {
        'label': '满族',
        'value': 3
      }, {
        'label': '其他',
        'value': 99
      }],
      politicalStatusOptions: [{
        'label': '团员',
        'value': 1
      }, {
        'label': '群众',
        'value': 2
      }],
      currentUnitOptions: [{
        'label': '兴安职业技术学院',
        'value': 1
      }],
      provinceOptions: [{
        'label': '内蒙古',
        'value': 1
      }],
      cityOptions: [{
        'label': '兴安盟',
        'value': 1
      }],
    }
  },
  computed: {},
  watch: {},
  created() { console.log(123) },
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
  }
}

</script>
<style>
</style>
