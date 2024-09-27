<template>
  <h1 class="p-4 text-xl">
    KẾT QUẢ GIÁM SÁT THỰC HIỆN KẾ HOẠCH NGOẠI KIỂM / CHỈNH SỬA
  </h1>
  <div class="flex flex-col p-4">
    <div class="flex flex-row justify-between w-full">
      <div class="flex items-center space-x-1">
        <h2 class="p-2 font-bold text-lg">Chương trình: {{ program.name }} </h2>
        <a href="#" @click="openModalProgram"><PencilIcon  class="w-6 h-6 text-blue-500" /></a>
      </div>
      <button
        class="px-4 py-2 text-sm font-bold text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring focus:ring-blue-300"
      >
        Thêm Mã
      </button>
    </div> 
    <div v-for="(code, index) in program.program_codes">
      <div class="flex items-center space-x-2">
        <h3 class="p-2">Mã: {{ code.name }}</h3>
       
        <a href="#" @click="openModalProgramCode(code)"> <PencilIcon  class="w-5 h-5 text-blue-500" /></a>
      </div>
      <div class="flex flex-col ml-3">
        <div
          class="flex flex-row space-x-2 items-center"
          v-for="(reminder, idx) in code.program_code_reminders"
        >
          <div class="p-3">Mẫu: {{ reminder.sample }}</div>
          <div class="p-3">
            Ngày nhận:
            <VueDatePicker
              v-if="moment(reminder.date_of_receive).year() !== 1"
              v-model="reminder.date_of_receive"
              format="dd/MM/yyyy"
            ></VueDatePicker>
            <VueDatePicker
              v-else
              v-model="date"
              format="dd/MM/yyyy"
            ></VueDatePicker>
          </div>
          <div class="p-3">
            Ngày trả:
            <VueDatePicker
              v-if="moment(reminder.date_of_return).year() !== 1"
              v-model="reminder.date_of_return"
              format="dd/MM/yyyy"
            ></VueDatePicker>
            <VueDatePicker
              v-else
              v-model="date"
              format="dd/MM/yyyy"
            ></VueDatePicker>
          </div>
          <div class="flex items-center">
            <input v-if="reminder.status == 0" type="checkbox" id="checkbox" checked class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
            <input v-else type="checkbox" id="checkbox" @change="toggleCheckbox(reminder, 0)" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
            <label for="checkbox" class="ml-2 text-sm text-gray-700">N/A</label>
          </div>
          <div class="flex items-center">
            <input v-if="reminder.status == 1" type="checkbox" id="checkbox" checked class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
            <input v-else type="checkbox" id="checkbox" @change="toggleCheckbox(reminder, 1)" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
            <label for="checkbox" class="ml-2 text-sm text-gray-700">Đạt</label>
          </div>
          <div class="flex items-center">
            <input v-if="reminder.status == 2" type="checkbox" id="checkbox" checked class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
            <input v-else type="checkbox" id="checkbox" @change="toggleCheckbox(reminder, 2)" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
            <label for="checkbox" class="ml-2 text-sm text-gray-700">Không Đạt</label>
          </div>
          <div class="p-3" v-show="reminder.status == 0">
             % đạt
             <input
                  type="number"
                  min="1" max="100"
                  v-model="reminder.percent_passed"
                  required
                  class="block w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:ring-blue-300"
              />
          </div>
          <div v-if="idx === code.program_code_reminders.length - 1" class="flex">
            <a href="#" @click="addMore(code)" class="flex"><PlusCircleIcon  class="w-6 h-6 text-blue-500" /> Add more</a>
          </div>
        </div>
      </div>
    </div>
    <div class="m-5">
      <button
        class="float-right px-4 py-2 font-bold text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring focus:ring-blue-300"
      >
        Update
      </button>
    </div>
  </div>
  <transition name="fade">
    <div v-if="isOpenEditProgram" class="fixed inset-0 flex items-center justify-center z-50">
      <div class="fixed inset-0 bg-black opacity-50" @click="closeModalProgram"></div>
      <div class="bg-white rounded-lg overflow-hidden shadow-lg z-10 max-w-sm w-full mx-4">
        <div class="p-4">
          <h2 class="text-lg font-semibold">Chương trình</h2>
          <p class="mt-2">
            <input
                  type="text"
                  v-model="program.name"
                  required
                  class="block w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:ring-blue-300"
                  placeholder="Chương trình"
              />
          </p>
        </div>
        <div class="p-4 flex justify-end">
          <button @click="closeModalProgram" class="bg-gray-500 text-white px-4 py-2 rounded mr-2">
            Close
          </button>
          <button @click="saveModalProgram" class="bg-blue-500 text-white px-4 py-2 rounded">
            Save
          </button>
        </div>
      </div>
    </div>
  </transition>
  <transition name="fade">
    <div v-if="isOpenEditProgramCode" class="fixed inset-0 flex items-center justify-center z-50">
      <div class="fixed inset-0 bg-black opacity-50" @click="closeModalProgram"></div>
      <div class="bg-white rounded-lg overflow-hidden shadow-lg z-10 max-w-sm w-full mx-4">
        <div class="p-4">
          <h2 class="text-lg font-semibold">Mã Chương trình</h2>
          <p class="mt-2">
            <input
                  type="text"
                  v-model="selectedProgramCode.name"
                  required
                  class="block w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:ring-blue-300"
                  placeholder="Mã chương trình"
              />
          </p>
        </div>
        <div class="p-4 flex justify-end">
          <button @click="closeModalProgramCode" class="bg-gray-500 text-white px-4 py-2 rounded mr-2">
            Close
          </button>
          <button @click="saveModalProgramCode" class="bg-blue-500 text-white px-4 py-2 rounded">
            Save
          </button>
          
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
import { useFetch } from "nuxt/app";
import { ref } from "vue";
import moment from "moment";
import { format } from "date-fns";
import { PencilIcon, PlusCircleIcon } from '@heroicons/vue/24/solid'

import VueDatePicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
const config = useRuntimeConfig();
const apiBase = config.public.apiBase;

const route = useRoute();

const program = ref({});
const programPrev = ref({});
const date = ref(new Date());
const isOpenEditProgram = ref(false);
const isOpenEditProgramCode = ref(false);
const selectedProgramCode = ref({});
const selectedProgramCodePrev = ref({});

onMounted(async () => {
  fetchData(route.params.id);
});

const fetchData = async (id) => {
  const { data } = await useFetch(`${apiBase}/programs/${id}`);
  if (data.value) {
    program.value = data.value.data;
    programPrev.value = {...data.value.data};
  }
};

const getValidDate = (date) => {
  if (moment(date).year() === "0001") {
    return moment();
  }
  return date;
};

const formatDate = (date) => {
  const day = date.getDate();
  const month = date.getMonth() + 1;
  const year = date.getFullYear();

  return `${day}/${month}/${year}`;
};

const openModalProgram = () => isOpenEditProgram.value = !isOpenEditProgram.value;
const closeModalProgram = () => {
  isOpenEditProgram.value = false
  program.value = {...programPrev.value}
};

const saveModalProgram = () => {
  isOpenEditProgram.value = false
  programPrev.value = {...program.value}
};

const openModalProgramCode = (code) => {
  isOpenEditProgramCode.value = !isOpenEditProgramCode.value;
  selectedProgramCode.value = code;
  selectedProgramCodePrev.value = {...code};
};

const closeModalProgramCode = () => {
  selectedProgramCode.value.name = selectedProgramCodePrev.value.name;
  isOpenEditProgramCode.value = false;
};

const saveModalProgramCode = () => {
  isOpenEditProgramCode.value = false;
};

const toggleCheckbox = (reminder, value) => {
  reminder.status = value;
};

const addMore = (code) => {
  let sample = 1;
  if (code.program_code_reminders.length > 0) {
    sample = code.program_code_reminders[code.program_code_reminders.length - 1].sample + 1;
  }
  let obj = {
    "program_code_id": code.id,
    "sample": sample,
    "date_of_receive": moment(),
    "date_of_return": moment(),
    "is_default": true,
    "status": 0,
    "percent_passed": 0
  };
  code.program_code_reminders = [
    ...code.program_code_reminders, obj
  ]
}

</script>
