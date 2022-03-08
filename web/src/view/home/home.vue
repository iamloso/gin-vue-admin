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
            <el-form-item label="报名项目" prop="professionalName" @change="changeProfessionalName">
              <el-select
                v-model="formData.professionalName"
                placeholder="请选择报名"
                clearable
                :style="{width: '100%'}"
              >
                <el-option
                  v-for="(item, index) in professionalNameOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="姓名" prop="name">
              <el-input
                v-model="formData.name"
                placeholder="请输入姓名"
                clearable
                :style="{width: '100%'}"
              />
            </el-form-item>
            <el-form-item label="手机号" prop="phone">
              <el-input v-model="formData.phone" placeholder="请输入手机号" clearable :style="{width: '100%'}" />
            </el-form-item>
            <el-form-item label="证件类型" prop="IDType">
              <el-select v-model="formData.IDType" placeholder="请输入证件类型" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in IDTypeOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="证件号" prop="UID">
              <el-input v-model="formData.UID" placeholder="请输入证件号" clearable :style="{width: '100%'}" @change="getUserByUID"/>
            </el-form-item>
            <el-form-item label="照片" prop="name">
              <div class="gva-btn-list">
                <el-upload
                    :action="`${path}/fileUploadAndDownload/upload`"
                    :before-upload="checkFile"
                    :headers="{ 'x-token': userStore.token, 'UID': formData.UID, 'picType': 'userPic' }"
                    :on-error="uploadError"
                    :on-success="uploadSuccessUserPic"
                    :show-file-list="false"
                    class="upload-btn"
                >
                  <el-button size="mini" type="primary">点击上传</el-button>
                </el-upload>
                <span style="padding-left: 20px; color: red">Tips: 二寸、白底、免冠</span>
              </div>
            </el-form-item>
            <el-form-item label="性别" prop="sex">
              <el-select v-model="formData.sex" placeholder="请选择性别" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in sexOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="民族" prop="nation">
              <el-select v-model="formData.nation" placeholder="请选择民族" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in nationOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="省份" prop="province">
              <el-select v-model="formData.province" placeholder="请选择省份" clearable :style="{width: '100%'}" @change="genCity">
                <el-option
                  v-for="(item, index) in province"
                  :key="index"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="城市" prop="city">
              <el-select v-model="formData.city" placeholder="请选择城市" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in city"
                  :key="index"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="证明" prop="name">
              <div class="gva-btn-list">
                <el-upload
                    :action="`${path}/fileUploadAndDownload/upload`"
                    :before-upload="checkFile"
                    :headers="{ 'x-token': userStore.token, 'UID': formData.UID, 'picType': 'userCertify' }"
                    :on-error="uploadError"
                    :on-success="uploadSuccessUserCertify"
                    :show-file-list="false"
                    class="upload-btn"
                >
                  <el-button size="mini" type="primary">点击上传</el-button>
                </el-upload>
                <span style="padding-left: 20px; color: red">上传：学籍在线证明或工作证明及身份证</span>
              </div>
            </el-form-item>
            <el-form-item label="所在单位" prop="currentUnit">
              <el-select v-model="formData.currentUnit" placeholder="请选择所在单位" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in currentUnitOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="出生日期" prop="dateBirth">
              <el-date-picker
                v-model="formData.dateBirth"
                type="date"
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
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="考生来源" prop="source">
              <el-select v-model="formData.source" placeholder="请选择考生来源" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in sourceOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="证书领取" prop="receive">
              <el-select v-model="formData.receive" placeholder="请选择证书领取" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in receiveOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="职业" prop="work">
              <el-select v-model="formData.work" placeholder="请选择职业" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in workOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item v-if="showLevel()" label="年级" prop="level">
              <el-select v-model="formData.level" placeholder="请选择年级" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in levelOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item v-if="showLevel()" label="班级" prop="conditions">
              <el-select v-model="formData.conditions" placeholder="请选择班级" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in conditionOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="参加工作时间" prop="workDate">
              <el-input v-model="formData.workDate" placeholder="请输入参加工作时间" clearable :style="{width: '100%'}" />
            </el-form-item>
            <el-form-item label="专业年限" prop="workYear">
              <el-input v-model="formData.workYear" placeholder="请输入专业年限" clearable :style="{width: '100%'}" />
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
                  :value="item.label"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
            <!--            <el-form-item label="工种名称" prop="workType">-->
            <!--              <el-select v-model="formData.workType" placeholder="请选择工种名称" clearable :style="{width: '100%'}" />-->
            <!--            </el-form-item>-->
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
            <el-form-item label="原证书职业" prop="OCO">
              <el-input v-model="formData.OCO" placeholder="请输入原证书职业" clearable :style="{width: '100%'}" />
            </el-form-item>
            <el-form-item label="原证书等级" prop="OCL">
              <el-input v-model="formData.OCL" placeholder="请输入原证书等级" clearable :style="{width: '100%'}" />
            </el-form-item>
            <el-form-item label="原证书编号" prop="OCN">
              <el-input v-model="formData.OCN" placeholder="请输入原证书编号" clearable :style="{width: '100%'}" />
            </el-form-item>
            <el-form-item label="">
              <img style="float: left" src="./pay.png" width="180">
            </el-form-item>
            <el-form-item label="扫码支付" prop="isPay">
              <el-radio-group v-model="formData.isPay" size="medium">
                <el-radio
                  v-for="(item, index) in paymentOptions"
                  :key="index"
                  :label="item.value"
                  :value="item.label"
                  :disabled="item.disabled"
                >{{ item.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form ref="elForm" :model="formData" :rules="rules" size="small" label-width="100px">
              <el-form-item label="支付金额" prop="payAmount">
                <el-input v-model="formData.payAmount" placeholder="请输入金额支付支付金额" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
            </el-form>
            <el-form-item label="支付凭证" prop="name">
              <div class="gva-btn-list">
                <el-upload
                    :action="`${path}/fileUploadAndDownload/upload`"
                    :before-upload="checkFile"
                    :headers="{ 'x-token': userStore.token, 'UID': formData.UID, 'picType': 'userPay' }"
                    :on-error="uploadError"
                    :on-success="uploadSuccessUserPay"
                    :show-file-list="false"
                    class="upload-btn"
                >
                  <el-button size="mini" type="primary">点击上传</el-button>
                </el-upload>
                <span style="padding-left: 20px; color: red">上传: 微信支付成功后的截图</span>
              </div>
            </el-form-item>
            <el-form-item size="large">
              <el-button type="primary" @click="submitForm">提交</el-button>
              <el-button @click="resetForm">重置</el-button>
              <el-button type="primary" @click="printReport">打印报名表</el-button>
            </el-form-item>
          </el-scrollbar>
        </el-form>
      </el-main>
      <el-footer>版权所有© Copyright 2022-2035 北京诗码作集科技有线公司技术支持(电话：18513172626)</el-footer>
    </el-container>
  </div>
</template>

<script setup>
import { useUserStore } from '@/pinia/modules/user'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const path = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

const imageUrl = ref('')

const fullscreenLoading = ref(false)
const checkFile = (file) => {
  let UID
  UID = window.localStorage.getItem('UID')
  if (!UID) {
    ElMessage.error('请先填写身份证号！')
    return false
  }
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
const uploadSuccessUserPic = (res) => {
  fullscreenLoading.value = false
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    if (res.code === 0) {
      window.localStorage.setItem('userPicUrl', res.data.file.url)
      console.log( window.localStorage.getItem('userPicUrl'))
    }
  } else {
    ElMessage({
      type: 'warning',
      message: res.msg
    })
  }
}
const uploadSuccessUserCertify = (res) => {
  fullscreenLoading.value = false
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    if (res.code === 0) {
      window.localStorage.setItem('userCertifyUrl', res.data.file.url)
      console.log( window.localStorage.getItem('userCertifyUrl'))
    }
  } else {
    ElMessage({
      type: 'warning',
      message: res.msg
    })
  }
}
const uploadSuccessUserPay = (res) => {
  fullscreenLoading.value = false
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    if (res.code === 0) {
      window.localStorage.setItem('userPayUrl', res.data.file.url)
      console.log( window.localStorage.getItem('userPayUrl'))
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
</script>

<script>
import {createJyxUser, getJyxUserList, findJyxUser} from '@/api/jyxUser'
import { ElMessage } from 'element-plus'

export default {
  name: 'Home',
  components: {},
  props: [],
  data() {
    return {
      province: ['北京市', '天津市', '河北省', '山西省', '内蒙古自治区', '辽宁省', '吉林省', '黑龙江省', '上海市', '江苏省', '浙江省', '安徽省', '福建省', '江西省', '山东省', '河南省', '湖北省', '湖南省', '广东省', '广西壮族自治区', '海南省', '重庆市', '四川省', '贵州省', '云南省', '西藏自治区', '陕西省', '甘肃省', '青海省', '宁夏回族自治区', '新疆维吾尔自治区', '新疆生产建设兵团'],
      provinceCity: {
        '北京市': ['东城区', '西城区', '朝阳区', '丰台区', '石景山区', '海淀区', '门头沟区', '房山区', '通州区', '顺义区', '昌平区', '大兴区', '怀柔区', '平谷区', '密云区', '延庆区'],
        '天津市': ['和平区', '河东区', '河西区', '南开区', '河北区', '红桥区', '东丽区', '西青区', '津南区', '北辰区', '武清区', '宝坻区', '滨海新区', '宁河区', '静海区', '蓟州区'],
        '河北省': ['石家庄市', '唐山市', '秦皇岛市', '邯郸市', '邢台市', '保定市', '张家口市', '承德市', '沧州市', '廊坊市', '衡水市'],
        '山西省': ['太原市', '大同市', '阳泉市', '长治市', '晋城市', '朔州市', '晋中市', '运城市', '忻州市', '临汾市', '吕梁市'],
        '内蒙古自治区': ['呼和浩特市', '包头市', '乌海市', '赤峰市', '通辽市', '鄂尔多斯市', '呼伦贝尔市', '巴彦淖尔市', '乌兰察布市', '兴安盟', '锡林郭勒盟', '阿拉善盟'],
        '辽宁省': ['沈阳市', '大连市', '鞍山市', '抚顺市', '本溪市', '丹东市', '锦州市', '营口市', '阜新市', '辽阳市', '盘锦市', '铁岭市', '朝阳市', '葫芦岛市'],
        '吉林省': ['长春市', '吉林市', '四平市', '辽源市', '通化市', '白山市', '松原市', '白城市', '延边朝鲜族自治州'],
        '黑龙江省': ['哈尔滨市', '齐齐哈尔市', '鸡西市', '鹤岗市', '双鸭山市', '大庆市', '伊春市', '佳木斯市', '七台河市', '牡丹江市', '黑河市', '绥化市', '大兴安岭地区'],
        '上海市': ['黄浦区', '徐汇区', '长宁区', '静安区', '普陀区', '虹口区', '杨浦区', '闵行区', '宝山区', '嘉定区', '浦东新区', '金山区', '松江区', '青浦区', '奉贤区', '崇明区'],
        '江苏省': ['南京市', '无锡市', '徐州市', '常州市', '苏州市', '南通市', '连云港市', '淮安市', '盐城市', '扬州市', '镇江市', '泰州市', '宿迁市'],
        '浙江省': ['杭州市', '宁波市', '温州市', '嘉兴市', '湖州市', '绍兴市', '金华市', '衢州市', '舟山市', '台州市', '丽水市'],
        '安徽省': ['合肥市', '芜湖市', '蚌埠市', '淮南市', '马鞍山市', '淮北市', '铜陵市', '安庆市', '黄山市', '滁州市', '阜阳市', '宿州市', '六安市', '亳州市', '池州市', '宣城市'],
        '福建省': ['福州市', '厦门市', '莆田市', '三明市', '泉州市', '漳州市', '南平市', '龙岩市', '宁德市'],
        '江西省': ['南昌市', '景德镇市', '萍乡市', '九江市', '新余市', '鹰潭市', '赣州市', '吉安市', '宜春市', '抚州市', '上饶市'],
        '山东省': ['济南市', '青岛市', '淄博市', '枣庄市', '东营市', '烟台市', '潍坊市', '济宁市', '泰安市', '威海市', '日照市', '临沂市', '德州市', '聊城市', '滨州市', '菏泽市'],
        '河南省': ['郑州市', '开封市', '洛阳市', '平顶山市', '安阳市', '鹤壁市', '新乡市', '焦作市', '濮阳市', '许昌市', '漯河市', '三门峡市', '南阳市', '商丘市', '信阳市', '周口市', '驻马店市', '济源市'],
        '湖北省': ['武汉市', '黄石市', '十堰市', '宜昌市', '襄阳市', '鄂州市', '荆门市', '孝感市', '荆州市', '黄冈市', '咸宁市', '随州市', '恩施土家族苗族自治州', '仙桃市', '潜江市', '天门市', '神农架林区'],
        '湖南省': ['长沙市', '株洲市', '湘潭市', '衡阳市', '邵阳市', '岳阳市', '常德市', '张家界市', '益阳市', '郴州市', '永州市', '怀化市', '娄底市', '湘西土家族苗族自治州'],
        '广东省': ['广州市', '韶关市', '深圳市', '珠海市', '汕头市', '佛山市', '江门市', '湛江市', '茂名市', '肇庆市', '惠州市', '梅州市', '汕尾市', '河源市', '阳江市', '清远市', '东莞市', '中山市', '潮州市', '揭阳市', '云浮市'],
        '广西壮族自治区': ['南宁市', '柳州市', '桂林市', '梧州市', '北海市', '防城港市', '钦州市', '贵港市', '玉林市', '百色市', '贺州市', '河池市', '来宾市', '崇左市'],
        '海南省': ['海口市', '三亚市', '三沙市', '儋州市', '五指山市', '琼海市', '文昌市', '万宁市', '东方市', '定安县', '屯昌县', '澄迈县', '临高县', '白沙黎族自治县', '昌江黎族自治县', '乐东黎族自治县', '陵水黎族自治县', '保亭黎族苗族自治县', '琼中黎族苗族自治县'],
        '重庆市': ['万州区', '涪陵区', '渝中区', '大渡口区', '江北区', '沙坪坝区', '九龙坡区', '南岸区', '北碚区', '綦江区', '大足区', '渝北区', '巴南区', '黔江区', '长寿区', '江津区', '合川区', '永川区', '南川区', '璧山区', '铜梁 区', '潼南区', '荣昌区', '开州区', '梁平区', '武隆区', '城口县', '丰都县', '垫江县', '忠县', '云阳县', '奉节县', '巫山县', '巫溪县', '石柱土家族自治县', '秀山土家族苗族自治县', '酉阳土家族苗族自治县', '彭水苗族土家族自治县'],
        '四川省': ['成都市', '自贡市', '攀枝花市', '泸州市', '德阳市', '绵阳市', '广元市', '遂宁市', '内江市', '乐山市', '南充市', '眉山市', '宜宾市', '广安市', '达州市', '雅安市', '巴中市', '资阳市', '阿坝藏族羌族自治州', '甘孜藏族自治州', '凉山彝族自治州'],
        '贵州省': ['贵阳市', '六盘水市', '遵义市', '安顺市', '毕节市', '铜仁市', '黔西南布依族苗族自治州', '黔东南苗族侗族自治州', '黔南布依族苗族自治州'],
        '云南省': ['昆明市', '曲靖市', '玉溪市', '保山市', '昭通市', '丽江市', '普洱市', '临沧市', '楚雄彝族自治州', '红河哈尼族彝族自治州', '文山壮族苗族自治州', '西双版纳傣族自治州', '大理白族自治州', '德宏傣族景颇族自治州', '怒江傈僳族自治州', '迪庆藏族自治州'],
        '西藏自治区': ['拉萨市', '日喀则市', '昌都市', '林芝市', '山南市', '那曲市', '阿里地区'],
        '陕西省': ['西安市', '铜川市', '宝鸡市', '咸阳市', '渭南市', '延安市', '汉中市', '榆林市', '安康市', '商洛市'],
        '甘肃省': ['兰州市', '嘉峪关市', '金昌市', '白银市', '天水市', '武威市', '张掖市', '平凉市', '酒泉市', '庆阳市', '定西市', '陇南市', '临夏回族自治州', '甘南藏族自治州'],
        '青海省': ['西宁市', '海东市', '海北藏族自治州', '黄南藏族自治州', '海南藏族自治州', '果洛藏族自治州', '玉树藏族自治州', '海西蒙古族藏族自治州'],
        '宁夏回族自治区': ['银川市', '石嘴山市', '吴忠市', '固原市', '中卫市'],
        '新疆维吾尔自治区': ['乌鲁木齐市', '克拉玛依市', '吐鲁番市', '哈密市', '昌吉回族自治州', '博尔塔拉蒙古自治州', '巴音郭楞蒙古自治州', '阿克苏地区', '克孜勒苏柯尔克孜自治州', '喀什地区', '和田地区', '伊犁哈萨克自治州', '塔城地区', '阿勒泰地区', '石河子市', '阿拉尔市', '图木舒克市', '五家渠市', '铁门关市'],
        '新疆生产建设兵团': ['乌鲁木齐市', '克拉玛依市', '吐鲁番市', '哈密市', '昌吉回族自治州', '博尔塔拉蒙古自治州', '巴音郭楞蒙古自治州', '阿克苏地区', '克孜勒苏柯尔克孜自治州', '喀什地区', '和田地区', '伊犁哈萨克自治州', '塔城地区', '阿勒泰地区', '石河子市', '阿拉尔市', '图木舒克市', '五家渠市', '铁门关市'],
      },
      city: [],
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
      sourceOptions: [{
        'label': '机关事业单位',
        'value': 1
      }, {
        'label': '职业技术学院',
        'value': 2
      }],
      receiveOptions: [{
        'label': '自取',
        'value': 1
      }],
      workOptions: [{
        'label': '学生',
        'value': 1
      }, {
        'label': '社会人员',
        'value': 2
      }],
      professionalNameOptions: [{
        'label': '保育员',
        'value': 1,
      }, {
        'label': '普通话',
        'value': 2
      }, { 'label': '1+X幼儿照护',
        'value': 3
      }],
      levelOptions: [{
        'label': '2020',
        'value': 1
      }, {
        'label': '2021',
        'value': 2
      }, {
        'label': '2022',
        'value': 3
      }],
      conditionOptions: [{
        'label': '1班',
        'value': 1
      }, {
        'label': '2班',
        'value': 2
      }, {
        'label': '3班',
        'value': 3
      }, {
        'label': '4班',
        'value': 4
      }, {
        'label': '5班',
        'value': 5
      }, {
        'label': '6班',
        'value': 6
      }, {
        'label': '7班',
        'value': 7
      }, {
        'label': '8班',
        'value': 8
      }, {
        'label': '9班',
        'value': 9
      }, {
        'label': '10班',
        'value': 10
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
        'label': '回族',
        'value': 4
      }, {
        'label': '藏族',
        'value': 5
      }, {
        'label': '维吾尔族',
        'value': 6
      }, {
        'label': '苗族',
        'value': 7
      }, {
        'label': '彝族',
        'value': 8
      }, {
        'label': '壮族',
        'value': 9
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
      paymentOptions: [{
        'label': '已支付',
        'value': '已支付'
      }, {
        'label': '未支付',
        'value': '未支付'
      }],
    }
  },
  computed: {},
  watch: {
    // 'formData.UID':function (data){
    //   window.localStorage.setItem('UID', data)
    // },
  },

  async created() {
    console.log(window.localStorage.getItem('UID'))
    if (window.localStorage.getItem('UID')) {
      this.formData.UID = window.localStorage.getItem('UID')
      const res = await findJyxUser({ UID: this.formData.UID})
      if (res.code === 0 && res.data.rejyxUser.ID > 0) {
        console.log(res.data.rejyxUser)
        this.formData = res.data.rejyxUser
        // await this.$router.push({ name: 'UploadFile', params: { name: this.formData.name, UID: this.formData.UID }})
      }
    }

  },
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(async valid => {
        if (!valid) return
        // TODO 提交表单
        this.formData.userPic = window.localStorage.getItem('userPicUrl')
        this.formData.userCertify = window.localStorage.getItem('userCertifyUrl')
        this.formData.userPay = window.localStorage.getItem('userPayUrl')
        console.log(this.formData.userPay)
        this.formData.UID = this.formData.UID.toUpperCase()
        const res = await createJyxUser(JSON.stringify(this.formData))
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '创建/更改成功'
          })
          window.localStorage.removeItem('UID')
          window.localStorage.removeItem('userPicUrl')
          window.localStorage.removeItem('userCertifyUrl')
          window.localStorage.removeItem('userPayUrl')
          //localStorage.clear()
          window.localStorage.setItem('UID', this.formData.UID)
          //this.formData = null
          // await this.$router.push({ name: 'UploadFile', params: { name: this.formData.name, UID: this.formData.UID }})
        }
      })
    },
    changeProfessionalName() {
       window.localStorage.removeItem('userCertifyUrl')
       window.localStorage.removeItem('userPayUrl')
       window.localStorage.removeItem('userPicUrl')
       this.formData.userCertify = ""
       this.formData.userPic = ""
       this.formData.userPay = ""
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
    genCity() {
      this.city = []
      this.formData.city = undefined
      for (const key in this.provinceCity) {
        if (key === this.formData.province) {
          this.city = this.provinceCity[key]
        }
      }
    },
    showLevel() {
      if (this.formData.professionalName === '普通话') {
        return true
      } else {
        return false
      }
    },
    async getUserByUID() {
      console.log(this.formData.UID)
      this.formData.UID = this.formData.UID.toUpperCase()
      window.localStorage.setItem('UID', this.formData.UID)
      const res = await findJyxUser({ UID: this.formData.UID})
      if (res.code === 0 && res.data.rejyxUser.ID > 0) {
        console.log(res.data.rejyxUser)
        this.formData = res.data.rejyxUser
        // await this.$router.push({ name: 'UploadFile', params: { name: this.formData.name, UID: this.formData.UID }})
      }
    },
    async printReport() {
      // let UID
      // UID = window.localStorage.getItem('UID')
      // if (!UID) {
      //   ElMessage.error('请先填写身份证号！')
      //   return false
      // }
      if (!this.formData.UID) {
        ElMessage.error('请先填写身份证号！')
        return false
      }
      await this.$router.push({name: 'Print'})
    },
  }
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
