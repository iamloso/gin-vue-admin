<template>
  <div v-loading.fullscreen.lock="fullscreenLoading" class="common-layout">
    <el-container>
      <el-header>
        <img style="float: left" src="./logo.png">
        <span style="float: left; font-size: 25px">师范教育系在线报名</span>
      </el-header>
      <div>
        <el-main>
          <el-scrollbar height="600px">
            <div class="gva-btn-list">
              <el-upload
                :action="`${path}/fileUploadAndDownload/upload`"
                :before-upload="checkFile"
                :headers="{ 'x-token': userStore.token }"
                :on-error="uploadError"
                :on-success="uploadSuccess"
                :show-file-list="false"
                class="upload-btn"
              >
                <el-button size="mini" type="primary">头像上传</el-button>
              </el-upload>
              <upload-image
                v-model:imageUrl="imageUrl"
                :file-size="512"
                :max-w-h="1080"
                class="upload-btn"
                @on-success="getTableData"
              />
            </div>

            <el-table :data="tableData">
              <el-table-column align="left" label="预览" width="100">
                <template #default="scope">
                  <CustomPic pic-type="file" :pic-src="scope.row.url" />
                </template>
              </el-table-column>
              <el-table-column align="left" label="日期" prop="UpdatedAt" width="180">
                <template #default="scope">
                  <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
                </template>
              </el-table-column>
              <el-table-column align="left" label="文件名" prop="name" width="180" />
              <el-table-column align="left" label="链接" prop="url" min-width="300" />
              <el-table-column align="left" label="标签" prop="tag" width="100">
                <template #default="scope">
                  <el-tag
                    :type="scope.row.tag === 'jpg' ? 'primary' : 'success'"
                    disable-transitions
                  >{{ scope.row.tag }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column align="left" label="操作" width="160">
                <template #default="scope">
                  <el-button size="small" icon="download" type="text" @click="downloadFile(scope.row)">下载</el-button>
                  <el-button size="small" icon="delete" type="text" @click="deleteFileFunc(scope.row)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
            <div class="gva-pagination">
              <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :style="{ float: 'right', padding: '20px' }"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
              />
            </div>
          </el-scrollbar>
        </el-main>
      </div>
      <el-footer>版权所有© Copyright 2022-2025 北京诗码作集科技有线公司技术支持(电话：18513172626)</el-footer>
    </el-container>
  </div>
</template>

<script setup>
import { getFileList, deleteFile } from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import { useUserStore } from '@/pinia/modules/user'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import { formatDate } from '@/utils/format'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const path = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

const imageUrl = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getFileList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const deleteFileFunc = async(row) => {
  ElMessageBox.confirm('此操作将永久文件, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteFile(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}

const fullscreenLoading = ref(false)
const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 0.5
  if (!isJPG && !isPng) {
    ElMessage.error('上传图片只能是 jpg或png 格式!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('未压缩未见上传图片大小不能超过 500KB，请使用压缩上传')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG) && isLt2M
}
const uploadSuccess = (res) => {
  fullscreenLoading.value = false
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    if (res.code === 0) {
      getTableData()
    }
  } else {
    ElMessage({
      type: 'warning',
      message: res.msg
    })
  }
}
const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}
const downloadFile = (row) => {
  if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    downloadImage(row.url, row.name)
  } else {
    downloadImage(path.value + row.url, row.name)
  }
}

</script>

<script>

export default {
  name: 'Upload',
  created() {
    console.log(this.$route.params.name)
  },
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
  line-height: 16px;
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
