<template>
  <v-container>
    <div class="wrapper_page">
      <h1 class="mb-10">ข้อมูลการตั้งค่า</h1>
      <v-row
        class="warpper-card mt-5 drop-shadow2 card"
        v-for="(item, i) in templates"
        :key="i"
      >
        <v-col style="text-align: start">
          <p class="mb-0" style="font-size: 20px">{{ item.name }}</p>
          <p class="mb-0" style="font-size: 30px; font-weight: bolder">
            {{ item.timestart + " : " + item.timeend }}
          </p>
          <b v-if="item.type == 0" class="mb-0" style="font-size: 15px">
            {{ checkday(item.daystart) + checkday(item.dayend) }}
          </b>
          <b v-if="item.type == 1" class="mb-0" style="font-size: 15px">
            {{ item.datestart + " ถึง " + item.dateend }}
          </b>
        </v-col>
        <v-col cols="3" class="d-flex">
          <v-btn icon color="#f9ae00" :to="`setting/${item.templateId}`">
            <v-icon size="35">mdi-pencil</v-icon>
          </v-btn>
          <v-btn icon color="red" @click="deleteitem(item)">
            <v-icon size="35">mdi-trash-can</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-btn
        block
        elevation="1"
        class="my-10"
        color="#526FFF"
        style="color: #fff; border-radius: 15px"
        to="setting/setting_templatetime"
        ><v-icon size="20">mdi-plus</v-icon>เพิ่ม</v-btn
      >
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      templates: [],
    };
  },
  mounted() {
    this.gettemplate();
  },
  methods: {
    gettemplate() {
      this.$axios
        .get(`${process.env.BASE_URL}/template/all`)
        .then((response) => {
          this.templates = response.data.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    checkday(day) {
      var days = [
        "อาทิตย์",
        "จันทร์",
        "อังคาร",
        "พุธ",
        "พฤหัสบดี",
        "ศุกร์",
        "เสาร์",
      ];
      return days[day];
    },
    deleteitem(item) {
      this.$axios
        .delete(`${process.env.BASE_URL}/template/delete`, {
          params: { id: item.templateId },
        })
        .then((response) => {
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
  align-content: center;
  p {
    margin: 0px;
  }
  .warpper-input {
    display: flex;
    flex-direction: row;
    padding: 10px;
    align-items: center;
    border-bottom: 1px solid #c4c4c4;
  }
}
</style>