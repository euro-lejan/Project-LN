<template>
  <div class="wrapper_page">
    <h1 class="mb-10">บันทึกหน่วยไฟฟ้า</h1>
    <el-date-picker
      style="width: 100%"
      v-model="search"
      type="year"
      placeholder="ปี"
      @change="getdata()"
    />
    <div class="card my-2">
      <highcharts
        v-if="showing"
        :options="chartOptions"
        style="border-radius: 15px"
      ></highcharts>
    </div>
    <!-- <v-col>
      <v-btn
        block
        elevation="2"
        color="#526FFF"
        style="color: #fff; border-radius: 15px"
        @click="dialog = !dialog"
        >เพิ่ม</v-btn
      >
    </v-col> -->
    <v-row>
      <v-col>
        <v-simple-table class="text-center">
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-center" style="width: 10%">ลำดับ</th>
                <th class="text-center">เดือน</th>
                <th class="text-center">หน่วย</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, i) in items" :key="i">
                <td>{{ i + 1 }}</td>
                <td>{{ checkmonth(new Date(item.date).getMonth()) }}</td>
                <td>{{ item.unit }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </v-row>
    <!-- <v-dialog v-model="dialog" width="500" class="dialog">
      <v-card style="padding: 20px" class="card">
        <p style="font-weight: bolder; font-size: 25px; text-align: center">
          เพิ่มหน่วยไฟฟ้า
        </p>
        <v-form ref="form" @submit.prevent="onsubmit">
          <v-menu
            ref="menu"
            v-model="dialog1"
            :close-on-content-click="false"
            :return-value.sync="addunit.date"
            transition="scale-transition"
            offset-y
            max-width="290px"
            min-width="auto"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                v-model="addunit.date"
                label="เลือกเดือน"
                readonly
                v-bind="attrs"
                v-on="on"
                solo
              ></v-text-field>
            </template>
            <v-date-picker
              v-model="addunit.date"
              type="month"
              no-title
              scrollable
            >
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="dialog1 = false">
                Cancel
              </v-btn>
              <v-btn
                text
                color="primary"
                @click="$refs.menu.save(addunit.date)"
              >
                OK
              </v-btn>
            </v-date-picker>
          </v-menu>
          <v-text-field
            type="number"
            label="ค่าวัตถ์"
            suffix="หน่วย"
            solo
            v-model="addunit.unit"
          ></v-text-field>

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
    </v-dialog> -->
  </div>
</template>

<script>
import { Chart } from "highcharts-vue";
import moment from "moment";
export default {
  components: {
    highcharts: Chart,
  },
  watch: {
    items(newval) {
      this.chartOptions.series.data = newval.map((e) => e.unit);
      this.chartOptions.title.text = `ปี ${moment(this.search).format("YYYY")}`;
      this.chartOptions.series.name = "ไฟฟ้าที่ใช้";
      this.chartOptions.xAxis.categories = newval.map((e) => moment(e.date).format("MMM"));
    },
    dialog(val) {
      if (!val) {
        this.addunit = {
          date: "",
          unit: 0,
        };
      }
    },
  },
  data() {
    return {
      search: moment().format("YYYY"),
      dialog: false,
      dialog1: false,
      addunit: {
        date: "",
        unit: 0,
      },
      items: [],
      unit: [],
      time: [],
      showing: true,
      chartOptions: {
        title: {
          text: "ค่าไฟฟ้าแต่ละปี",
          align: "center",
        },
        yAxis: {
          title: {
            text: "จำนวนหน่วย",
          },
        },
        chart: {
          text: "จำนวนหน่วย",
          height: 200,
        },
        xAxis: {
          categories: [
            "1",
            "2",
            "3",
            "4",
            "5",
            "6",
            "7",
            "8",
            "9",
            "10",
            "11",
            "12",
          ],
        },
        legend: {
          layout: "vertical",
          align: "right",
          verticalAlign: "middle",
        },

        series: {
          name: "ค่าไฟ(หน่วย)",
          data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        },

        responsive: {
          rules: [
            {
              condition: {
                maxWidth: 500,
              },
              chartOptions: {
                legend: {
                  layout: "horizontal",
                  align: "center",
                  verticalAlign: "bottom",
                },
              },
            },
          ],
        },
        colors: ["#3336FF"],
      },
    };
  },
  mounted() {
    this.getdata();
  },
  methods: {
    checkmonth(index){
      const month = ["January","February","March","April","May","June","July","August","September","October","November","December"];
      return month[index]
    },
    async getdata() {
      var datestart = moment(this.search);
      var dateend = moment(this.search).add(1, "y");
      await this.$axios
        .post(`${process.env.BASE_URL}/saveunit/all`, {
          datestart,
          dateend,
        })
        .then((response) => {
          this.items = response.data.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    // async onsubmit() {
    //   this.$axios
    //     .post(`${process.env.BASE_URL}/saveunit/add`, {
    //       ...this.addunit,
    //       unit: parseInt(this.addunit.unit),
    //     })
    //     .then((response) => {
    //       // this.items.push(response.data.data);
    //       location.reload();
    //       this.dialog = false;
    //     })
    //     .catch((err) => {
    //       console.log(err);
    //     });
    // },
  },
};
</script>

<style lang="scss" scoped>
</style>
