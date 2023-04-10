<template>
  <v-container>
    <div class="wrapper_page">
      <h1>สถานะไฟฟ้า</h1>
      <v-row
        class="warpper-card mt-5 drop-shadow2 card"
        v-for="(item, i) in items"
        :key="i"
      >
        <v-col style="text-align: start">
          <b class="mb-0"
            ><span>{{ item.name }}</span>
            <span>
              <v-btn
                class="m-0"
                color="orange"
                dark
                small
                @click="editchanal(item, i)"
              >
                <v-icon dark left> mdi-pencil-box-multiple </v-icon>
                แก้ไข
              </v-btn></span
            ></b
          >
          <p class="mb-0" style="font-size: 18px">
            รูปแบบ : {{ !item.templatestatus ? "กำหนดเอง" : "ตั้งเวลา" }}
          </p>
          <p v-if="item.templatestatus" class="mb-0" style="font-size: 15px">
            ชื่อรูปแบบ : {{ item.template?.name ? item.template.name : "-"  }}
          </p>
        </v-col>
        <v-switch
          :disabled="item.templatestatus"
          class="m-0"
          color="success"
          v-model="item.status"
          inset
          @change="onChange(item)"
        ></v-switch>
        <!--v-switch ประกาศ@change="onChange(ตัวแปรที่จะเรียกใช้ในฟังก์ชัน  methods onchang) -->
      </v-row>
    </div>
    <v-dialog v-model="dialog" width="500" class="dialog">
      <v-card style="padding: 20px" class="card">
        <p style="font-weight: bolder">แก้ไข</p>
        <v-form ref="form" @submit.prevent="onsubmit">
          <v-text-field solo v-model="detail.name" placeholder="ชื่อ" />
          <v-select
            :items="type"
            item-value="id"
            v-model="detail.templatestatus"
            item-text="name"
            placeholder="เลือกรูปแบบ"
            solo
          ></v-select>
          <v-select
            v-if="detail.templatestatus"
            :items="template"
            item-value="templateId"
            v-model="detail.templateId"
            item-text="name"
            placeholder="เลือกรูปแบบ"
            solo
          ></v-select>
          <v-btn
            block
            elevation="2"
            color="#526FFF"
            style="color: #fff; border-radius: 15px"
            type="submit"
            >บันทึก</v-btn
          >
        </v-form>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      items: [],
      template: [],
      type: [
        { id: false, name: "กำหนดเอง" },
        { id: true, name: "ตั้งเวลา" },
      ],
      dialog: false,
      detail: {
        name: "",
        type: null,
        templateId: null,
      },
    };
  },
  mounted() {
    this.getdata();
    this.gettemplate();
  },
  methods: {
    getdata() {
      this.$axios
        .get(`${process.env.BASE_URL}/channels/all`)
        .then((response) => {
          this.items = response.data.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    gettemplate() {
      this.$axios
        .get(`${process.env.BASE_URL}/template/all`)
        .then((response) => {
          this.template = response.data.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    onChange(value) {
      this.$axios
        .put(`${process.env.BASE_URL}/channels/changestatus`, {
          channelId: value.channelId,
          status: value.status,
        })
        .then((response) => {})
        .catch((err) => {
          console.log(err);
        });
    },
    editchanal(value) {
      this.detail = JSON.parse(JSON.stringify(value));
      this.dialog = !this.dialog;
    },
    async onsubmit() {
      console.log(this.detail);
      await this.$axios
        .post(`${process.env.BASE_URL}/channels/update`, {
          channelId: this.detail.channelId,
          name: this.detail.name,
          templateId: this.detail.templateId,
          templatestatus: this.detail.templatestatus,
        })
        .then((response) => {
          this.dialog = !this.dialog;
          location.reload();
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.wrapper_page {
  flex-direction: row;
  align-content: center;

  .warpper-card {
    padding: 10px 20px;
    background: #fff;
    border-radius: 15px;
    font-size: 20px;
    justify-content: space-around;
    span {
      display: inline-block;
      max-width: 120px;
      white-space: nowrap;
      overflow: hidden !important;
      text-overflow: ellipsis;
    }
  }
}
.card {
  text-align: center;
  p {
    font-size: 30px;
  }
}
</style>
