<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="通讯地址" prop="address" width="120" />
        <el-table-column align="left" label="城市" prop="city" width="120" />
        <el-table-column align="left" label="申报条件" prop="conditions" width="120" />
        <el-table-column align="left" label="所在单位" prop="currentUnit" width="120" />
        <el-table-column align="left" label="出生日期" prop="dateBirth" width="120" />
        <el-table-column align="left" label="文化程度" prop="eduLevel" width="120" />
        <el-table-column align="left" label="邮箱" prop="email" width="120" />
        <el-table-column align="left" label="简要经历" prop="experience" width="120" />
        <el-table-column align="left" label="证件类型" prop="IDType" width="120" />
        <el-table-column align="left" label="级别" prop="level" width="120" />
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="民族" prop="nation" width="120" />
        <el-table-column align="left" label="原证书等级" prop="OCL" width="120" />
        <el-table-column align="left" label="原证书编号" prop="OCN" width="120" />
        <el-table-column align="left" label="原证书职业" prop="OCO" width="120" />
        <el-table-column align="left" label="手机号" prop="phone" width="120" />
        <el-table-column align="left" label="户籍所在地" prop="place" width="120" />
        <el-table-column align="left" label="政治面貌" prop="politicalStatus" width="120" />
        <el-table-column align="left" label="邮寄地址" prop="postAddress" width="120" />
        <el-table-column align="left" label="邮政编码" prop="postalCode" width="120" />
        <el-table-column align="left" label="职业名称" prop="professionalName" width="120" />
        <el-table-column align="left" label="省份" prop="province" width="120" />
        <el-table-column align="left" label="证件领取" prop="receive" width="120" />
        <el-table-column align="left" label="学历证书编号" prop="serialNumber" width="120" />
        <el-table-column align="left" label="性别" prop="sex" width="120" />
        <el-table-column align="left" label="考生来源" prop="source" width="120" />
        <el-table-column align="left" label="身份证号" prop="UID" width="120" />
        <el-table-column align="left" label="从事专业" prop="work" width="120" />
        <el-table-column align="left" label="参加工作时间" prop="workDate" width="120" />
        <el-table-column align="left" label="工种名称" prop="workType" width="120" />
        <el-table-column align="left" label="专业年限" prop="workYear" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateJyxUserFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
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
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteJyxUser,
  deleteJyxUserByIds,
  updateJyxUser,
  findJyxUser,
  getJyxUserList
} from '@/api/jyxUser'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getJyxUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteJyxUserFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteJyxUserByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateJyxUserFunc = async(row) => {
  const res = await findJyxUser({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rejyxUser
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteJyxUserFunc = async(row) => {
  const res = await deleteJyxUser({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
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
  }
}
// 弹窗确定
const enterDialog = async() => {
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
    closeDialog()
    getTableData()
  }
}
</script>

<style>
</style>
