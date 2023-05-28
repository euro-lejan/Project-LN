<template>
  <div class="wrapper_page">
    <h1 class="mb-10">บันทึกหน่วยไฟฟ้า</h1>
    <v-row>
      <v-col cols="5">
        <el-select v-model="type" placeholder="Select">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option> </el-select
      ></v-col>
      <v-col>
        <el-date-picker
          v-if="type == 'month'"
          style="width: 100%"
          v-model="search"
          type="year"
          placeholder="ปี"
          @change="getdata()"
          format="yyyy"
          value-format="yyyy-MM-dd"
        />
        <el-date-picker
          v-else-if="type == 'day'"
          style="width: 100%"
          v-model="search"
          type="month"
          placeholder="เดือน"
          @change="getdata()"
          format="MM-yyyy"
          value-format="yyyy-MM-dd"
        />
        <el-date-picker
          v-else
          style="width: 100%"
          v-model="search"
          type="date"
          placeholder="วัน"
          @change="getdata()"
          format="yyyy-MM-dd"
          value-format="yyyy-MM-dd"
        />
      </v-col>
    </v-row>
    <div class="card my-2">
      <highcharts
        v-if="showing"
        :options="chartOptions"
        style="border-radius: 15px"
      ></highcharts>
    </div>
    <v-col>
      <el-input
        class="mb-2"
        type="number"
        v-model="cost_f"
        placeholder="กรอกค่าไฟ/หน่วย"
      />
      <v-btn
        block
        elevation="2"
        color="#526FFF"
        style="color: #fff; border-radius: 15px"
        @click="sum()"
        >คำนวณค่าไฟ</v-btn
      >
    </v-col>
    <v-row>
      <v-col>
        <v-simple-table class="text-center">
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-center">
                  {{
                    type == "month" ? "เดือน" : type == "day" ? "วัน" : "เวลา"
                  }}
                </th>
                <th class="text-center">หน่วย</th>
                <th class="text-center">ค่าไฟ</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, i) in items" :key="i">
                <td>
                  {{
                    type == "month"
                      ? checkmonth($moment(item.date).month())
                      : type == "day"
                      ? $moment(item.date).format("DD/MM/yyyy")
                      : $moment(item.date).subtract(7, "hour").format("HH:mm")
                  }}
                </td>
                <td>
                  {{ (item.unit/1000).toLocaleString() }}
                </td>
                <td>
                  {{ ((item.unit/1000) * cost).toLocaleString() }}
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import { Chart } from "highcharts-vue";
import moment from "moment";
import constants from "@/constants/variable";
export default {
  components: {
    highcharts: Chart,
  },
  watch: {
    items(newval) {
      this.chartOptions.series.data = newval.map((e) => e.unit);
      this.chartOptions.title.text =
        this.type == "month"
          ? `ปี ${moment(this.search).format("yyyy")}`
          : this.type == "day"
          ? `เดือน ${
              constants.TH_MONTH.FULL[moment(this.search).month() + 1] +
              " " +
              moment(this.search).format("yyyy")
            }`
          : `วัน ${moment(this.search).format("DD-MM-yyyy")}`;
      this.chartOptions.series.name = "ไฟฟ้าที่ใช้";
      this.chartOptions.xAxis.categories = newval.map((e) =>
        this.type == "month"
          ? constants.TH_MONTH.FULL[moment(e.date).month() + 1]
          : this.type == "day"
          ? moment(e.date).format("DD/MM")
          : moment(e.date).subtract(7, "hour").format("HH:mm")
      );
    },
    type() {
      this.search = moment(this.search).format("yyyy-01-01");
      this.getdata();
    },
  },
  data() {
    return {
      type: "month",
      options: [
        {
          value: "month",
          label: "รายปี",
        },
        {
          value: "day",
          label: "รายเดือน",
        },
        {
          value: "hour",
          label: "รายวัน",
        },
      ],
      search: moment().format("yyyy-01-01"),
      dialog: false,
      dialog1: false,
      addunit: {
        date: "",
        unit: 0,
      },
      cost_f: 3.4,
      cost: null,
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
          categories: [],
        },
        legend: {
          layout: "vertical",
          align: "right",
          verticalAlign: "middle",
        },

        series: {
          name: "ค่าไฟ(หน่วย)",
          data: [],
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
    this.cost = this.cost_f;
  },
  methods: {
    sum() {
      this.cost = this.cost_f;
    },
    checkmonth(index) {
      return constants.TH_MONTH.FULL[index + 1];
    },
    async getdata() {
      console.log(this.search);
      console.log(moment(this.search).format("yyyy-MM-DD"));
      var datestart = moment(this.search).format("yyyy-MM-DD");
      var dateend = moment(this.search)
        .add(1, this.type == "month" ? "y" : this.type == "day" ? "M" : "d")
        .format("yyyy-MM-DD");

      console.log({
        time: datestart,
        dateend,
        type: this.type,
      });
      await this.$axios
        .post(`${process.env.BASE_URL}/saveunit/all`, {
          datestart,
          dateend,
          type: this.type,
        })
        .then((response) => {
          this.items = response.data.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>

<style lang="scss" scoped></style>
