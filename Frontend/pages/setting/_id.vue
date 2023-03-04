<template>
  <div class="wrapper_page">
    <h1 class="mb-10">ตั้งค่าเวลาเปิด</h1>
    <v-row class="warpper-input mt-5">
      <v-form ref="form" @submit.prevent="onsubmit">
        <v-row class="warpper-input">
          <v-text-field
            v-model="time.name"
            solo
            placeholder="ชื่อ"
          ></v-text-field>
          <v-select
            :items="type"
            item-value="id"
            item-text="name"
            v-model="time.type"
            placeholder="เลือกรูปแบบ"
            solo
          ></v-select>
          <p v-if="time.type != 1">เลือกวัน</p>
          <v-row class="d-flex" style="align-items: baseline">
            <v-col>
              <v-select
                
                v-if="time.type != 1"
                :items="items"
                item-value="id"
                item-text="name"
                v-model="time.daystart"
                solo
                placeholder="วันเริ่มต้น"
              ></v-select>
            </v-col>
            <p>{{ " - " }}</p>
            <v-col>
              <v-select
               
                v-if="time.type != 1"
                :items="items"
                item-value="id"
                item-text="name"
                v-model="time.dayend"
                solo
                placeholder="วันสิ้นสุด"
              ></v-select>
            </v-col>
          </v-row>
          <p>เวลาเริ่มต้น</p>
          <v-dialog
            v-if="time.type == 1"
            ref="dialog1"
            v-model="datestart"
            :return-value.sync="time.datestart"
            persistent
            width="290px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                v-model="time.datestart"
                label="เลือกวันที่"
                readonly
                type="datetime"
                solo
                v-bind="attrs"
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker v-model="time.datestart" scrollable>
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="datestart = false">
                Cancel
              </v-btn>
              <v-btn
                text
                color="primary"
                @click="$refs.dialog1.save(time.datestart)"
              >
                OK
              </v-btn>
            </v-date-picker>
          </v-dialog>
          <v-text-field
            v-model="time.timestart"
            solo
            type="time"
            placeholder="เวลา HH:MM"
          ></v-text-field>
          <p>เวลาสิ้นสุด</p>
          <v-dialog
            v-if="time.type == 1"
            ref="dialog"
            v-model="dateend"
            :return-value.sync="time.dateend"
            persistent
            width="290px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                v-model="time.dateend"
                label="เลือกวันที่"
                readonly
                solo
                v-bind="attrs"
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker v-model="time.dateend" scrollable>
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="dateend = false">
                Cancel
              </v-btn>
              <v-btn
                text
                color="primary"
                @click="$refs.dialog.save(time.dateend)"
              >
                OK
              </v-btn>
            </v-date-picker>
          </v-dialog>
          <v-text-field
            v-model="time.timeend"
            solo
            type="time"
            placeholder="เวลา HH:MM"
          ></v-text-field>
        </v-row>
        <v-col class="mb-5 wrapper-btn">
          <v-btn
            block
            elevation="2"
            color="#526FFF"
            style="color: #fff; border-radius: 15px"
            type="submit"
            >บันทึก</v-btn
          >
        </v-col>
      </v-form>
    </v-row>
  </div>
</template>
<script>
//import VueDayjs from 'vue-dayjs-plugin';
export default {
  data() {
    return {
      datestart: false,
      dateend: false,
      time: {
        type: 0,
        daystart: 1,
        dayend: 1,
        datestart: null,
        dateend: null,

        timestart: null,
        timeend: null,
      },
      items: [
        { id: 0, name: "อาทิตย์" },
        { id: 1, name: "จันทร์" },
        { id: 2, name: "อังคาร" },
        { id: 3, name: "พุธ" },
        { id: 4, name: "พฤหัสบดี" },
        { id: 5, name: "ศุกร์" },
        { id: 6, name: "เสาร์" },
      ],
      type: [
        { id: 0, name: "รายสัปดาห์" },
        { id: 1, name: "กำหนดเอง" },
      ],
    };
  },
  mounted() {
    this.$axios
      .get(`${process.env.BASE_URL}/template/detail`, {
        params: { id: this.$route.params.id },
      })
      .then((response) => {
        this.time = response.data.data;
      })
      .catch((err) => {
        console.log(err);
      });
  },
  methods: {
    // FUNTION ONSUBMIT
    onsubmit() {
      this.$axios
        .put(`${process.env.BASE_URL}/template/update`, this.time)
        .then((response) => {
          this.$router.go(-1);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    // END FUNTION ONSUBMIT
  },
};
</script>

<style lang="scss" scoped>
.wrapper_page {
  align-content: center;

  .warpper-input {
    display: flex;
    flex-direction: column;
    padding: 0px 20px;
    justify-content: center;
    text-align: left;
  }
  i {
    font-size: 40px !important;
  }
}
.content-date {
  height: 24px;
  width: 50px;
  background-color: red;
}
</style>
