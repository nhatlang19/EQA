<template>
    <h1 class="p-4 text-xl">KẾT QUẢ GIÁM SÁT THỰC HIỆN KẾ HOẠCH NGOẠI KIỂM</h1>
    <div class="text-right mr-2">
      <button
        @click="exportHandler()"
        class="px-4 py-2 text-sm font-bold text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring focus:ring-blue-300"
      >
        Export
      </button>
    </div>
    <div class="flex">
      <table class="m-2 table-auto w-full border-collapse border border-slate-500 text-md">
        <thead>
          <tr>
            <th class="w-16 border border-slate-600 p-2">STT</th>
            <th class="w-16 border border-slate-600 p-2 break-words" >Nội dung kế hoạch</th>
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
            <th class="w-16 border border-slate-600 p-2 break-words"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(program, index) in items">
            <td class="border border-slate-600 p-2 text-center">{{ index + 1 }}</td>
            <td class="border border-slate-600 p-2">{{ program.name }}</td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(1, index) }'><Tooltip :text="getByPassed(1, index)">{{ getByMonth(1, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(2, index) }'><Tooltip :text="getByPassed(2, index)">{{ getByMonth(2, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(3, index) }'><Tooltip :text="getByPassed(3, index)">{{ getByMonth(3, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(4, index) }'><Tooltip :text="getByPassed(4, index)">{{ getByMonth(4, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(5, index) }'><Tooltip :text="getByPassed(5, index)">{{ getByMonth(5, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(6, index) }'><Tooltip :text="getByPassed(6, index)">{{ getByMonth(6, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(7, index) }'><Tooltip :text="getByPassed(7, index)">{{ getByMonth(7, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(8, index) }'><Tooltip :text="getByPassed(8, index)">{{ getByMonth(8, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(9, index) }'><Tooltip :text="getByPassed(9, index)">{{ getByMonth(9, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(10, index) }'><Tooltip :text="getByPassed(10, index)">{{ getByMonth(10, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(11, index) }'><Tooltip :text="getByPassed(11, index)">{{ getByMonth(11, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words" :class='{"text-red-600":  isFuture(12, index) }'><Tooltip :text="getByPassed(12, index)">{{ getByMonth(12, index) }}</Tooltip></td>
            <td class="border border-slate-600 text-center p-2 whitespace-normal break-words"><nuxt-link :to="`/programs/${program.id}`" class="text-blue-600 hover:text-blue-600">Edit</nuxt-link></td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script setup>
  import { useFetch } from "nuxt/app";
  import { ref } from "vue";
  import Tooltip from '~/components/Tooltip.vue';
  import moment from "moment";
  const config = useRuntimeConfig();
  const apiBase = config.public.apiBase;
  
  const items = ref([]);
  const programSampleList = [];
  const programNameList = [];
  const programDate = [];
  const programStatus = [];
  const programPercent = [];
  
  onMounted(async () => {
    fetchData();
  });
  
  const fetchData = async () => {
    const { data: programs } = await useFetch(`${apiBase}/programs`);
    if (programs.value) {
      items.value = programs.value.data;
      convertData();
    }
  }

  const exportHandler = async () => {
    const response = await fetch(`${apiBase}/programs/export`, {
      method: 'GET',
    });
    if (response.ok) {
      const blob = await response.blob()
      const link = document.createElement('a')
      link.href = URL.createObjectURL(blob)
      link.download = 'program.xlsx'
      link.click()
      URL.revokeObjectURL(link.href)
    } else {
      console.error('Failed to download file')
    }
  }
  
  const convertData = () => {
    for (let index = 0; index < items.value.length; index++) {
      let programCodes = items.value[index].program_codes;
      for (const idx in programCodes) {
          let name = programCodes[idx].name;
          let reminders =  programCodes[idx].program_code_details;
          
          for (const reminderIndex in reminders) {
              let reminder = reminders[reminderIndex];
              let monthReminder = moment(reminder.date_of_return).month() + 1;
              if (!programSampleList[`${index} - ${monthReminder}`]) {
                programSampleList[`${index} - ${monthReminder}`] = []
              }
              programSampleList[`${index} - ${monthReminder}`].push(formatSample(reminder.sample))
              programNameList[`${index} - ${monthReminder}`] = name
              programDate[`${index} - ${monthReminder}`] = reminder.date_of_return

              programStatus[`${index} - ${monthReminder} - ${formatSample(reminder.sample)}`] = reminder.status
              programPercent[`${index} - ${monthReminder} - ${formatSample(reminder.sample)}`] = reminder.percent_passed
          }
      }
    }
  }

  const getByPassed = (month, index) => {
    let result = '';
    
    let name = programNameList[`${index} - ${month}`];
    let list = programSampleList[`${index} - ${month}`];
    if (list && list.length) {
      for (let sample of list) {
        let status = programStatus[`${index} - ${month} - ${sample}`];
        let percent = programPercent[`${index} - ${month} - ${sample}`];
        if (status == 1) {
          result += `${name}-${sample}: đạt ${percent}% <br />`;
        }
        if (status == 2) {
          result += `${name}-${sample}: không đạt <br />`;
        }
      }
    }
    
    return result;
  }
  
  const getByMonth = (month, index) => {
    let name = programNameList[`${index} - ${month}`];
    let list = programSampleList[`${index} - ${month}`]
    return formatList(list, name);
  }
  
  const isFuture = (month, index) => {
    let date = programDate[`${index} - ${month}`];
    return moment(date) > moment();
  }
  
  const formatList = (list, name) => {
    if (list && list.length) {
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
  