<template>
  <h1 class="p-4 text-xl">KẾT QUẢ GIÁM SÁT THỰC HIỆN KẾ HOẠCH NGOẠI KIỂM</h1>
  <div class="flex">
    <table class="m-2 table-auto w-full border-collapse border border-slate-500 text-md">
      <thead>
        <tr>
          <th class="w-16 border border-slate-600 p-2">STT</th>
          <th class="w-16 border border-slate-600 p-2 break-words">Nội dung kế hoạch</th>
          <th class="w-16 border border-slate-600 p-4 break-words ">Tháng 01</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 02</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 03</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 04</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 05</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 06</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 07</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 08</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 09</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 10</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 11</th>
          <th class="w-16 border border-slate-600 p-4 break-words">Tháng 12</th>
          <th class="w-16 border border-slate-600 p-2 break-words">Kết quả</th>
          <th class="w-16 border border-slate-600 p-2 break-words"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(program, index) in items">
          <td class="border border-slate-600 p-2 text-center">{{ index + 1 }}</td>
          <td class="border border-slate-600 p-2">{{ program.name }}</td>
          <td class="border bg-red border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(1, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(2, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(3, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(4, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(5, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(6, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(7, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(8, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(9, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(10, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(11, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">{{ groupByMonth(12, index) }}</td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words"></td>
          <td class="border border-slate-600 text-center p-2 whitespace-normal overflow-hidden break-words">Edit</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { useFetch } from "nuxt/app";
import { ref } from "vue";
import moment from "moment";
const config = useRuntimeConfig();
const apiBase = config.public.apiBase;
const { data: programs } = await useFetch(`${apiBase}/programs`);

const items = ref([]);
if (programs.value) {
  items.value = programs.value.data; // Assign the fetched data to a reactive reference if needed
}

const isFuture = (month) => {
    return (month > moment().month() + 1);
}

const groupByMonth = (month, index) => {
    let programCodes = items.value[index].program_codes;
    for (const idx in programCodes) {
        let name = programCodes[idx].name;
        let reminders =  programCodes[idx].program_code_reminders;
        let list = [];
        for (const reminderIndex in reminders) {
            let reminder = reminders[reminderIndex];
            let monthReminder = moment(reminder.date_of_return).month() + 1;
            if (monthReminder === month) {
                list.push(formatSample(reminder.sample));
            }
        }
        return formatList(list, name);
    }
    return '';
};

const formatList = (list, name) => {
    let res = [];
    if (list.length) {
        if (list.length == 1) {
            return `${name} ${list[0]}`;
        }
        return `${name} ${list[0]}-${list[list.length - 1]}`;
    }
    return '';
}

const formatSample = (sample) => {
    if (sample < 10) {
        return `0${sample}`;
    }
    return `${sample}`;
}
// definePageMeta({
//     middleware: 'auth'
// });
</script>
