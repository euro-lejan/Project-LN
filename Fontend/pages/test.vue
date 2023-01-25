<template>
  <v-form ref="form" v-model="valid" lazy-validation>
    <v-text-field
      v-model="name"
      :counter="10"
      :rules="nameRules"
      label="Name"
      required
    ></v-text-field>

    <v-btn color="success" class="mr-4" @click="validate">
      SUBMIT
    </v-btn>
  </v-form>
</template>


<script>
import axios from "axios";
export default {
  data() {
    return {
      time: {
        type: null,
        day: null,
        strat_date: "00/00/0000",
        end_date: "00/00/0000",
        strat_time: "00:00",
        end_time: "00:00",
      },
      items: [
        "อาทิตย์",
        "จันทร์",
        "อังคาร",
        "พุธ",
        "พฤหัสบดี",
        "ศุกร์",
        "เสาร์",
      ],
      type: [
        { id: 1, name: "รายสัปดาห์" },
        { id: 2, name: "กำหนดเอง" },
      ],
    };
  },
  methods: {
    // FUNTION ONSUBMIT
    onsubmit() {
      axios
        .post("http://127.0.0.1:4003/savetime", {
          type: this.time.type,
          day: this.time.day,
          starttime: this.time.strat_time,
          endtime: this.time.end_time,
        })
        .then((response) => {
          console.log(response.data);
        });
    },
    // END FUNTION ONSUBMIT
    validate() {
      //alert("HI");

      //console.log(this.time.type);
      //console.log(this.time.day);
      //console.log(this.time.strat_time);
      //console.log(this.time.end_time);

      axios
        .post("http://127.0.0.1:4003/savetime", {
            type: '2',      // Type
            daystart: '2022-10-28',   // วันเริ่มต้น
            starttime: '00:00',  // เวลาเริ่มต้น
            dayend: '00:30',     // วันสิ้นสุด
            endtime: '2022-10-29'     // เวลาสิ้นสุด
        })
        .then((response) => {

          console.log(response.data);
          
        });
    },
  },
};
</script>
