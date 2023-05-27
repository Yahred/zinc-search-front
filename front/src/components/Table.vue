<script setup lang="ts">
import { ref, onMounted, watchPostEffect, withDefaults } from 'vue'

import { type TableHeader } from '../interfaces'
import Spinner from './Spinner.vue';

import axios from '@/config/axios';

interface TableProps {
    data?: any[]
    definition: TableHeader[],
    uri?: string,
    search?: string,
    pageSizeOptions: number[],
}

const props = withDefaults(defineProps<TableProps>(), {
    pageSizeOptions: () => [5, 10, 15],
    data: () => [],
})

const emit = defineEmits(['click']);

const data = ref<any[]>(props.data);

const pageSize = ref<number>(5);
const currentPage = ref<number>(1);
const isLoading = ref<boolean>(false);


function emitClick(item: any) {
    emit('click', item);
}

async function fetchData() {
    if (!props.uri) return;

    const response = await axios.get<unknown, any[]>(props.uri, {
        params: {
            pageSize: pageSize.value,
            page: currentPage.value,
            q: props.search
        }
    });
    setTimeout(() => isLoading.value = false, 200);
    data.value = response;
}


function nextPage() {
    if (data.value.length < pageSize.value) return;

    currentPage.value++;
    isLoading.value = true;

    fetchData();
}

function prevPage() {
    if (currentPage.value === 1) return;

    currentPage.value--;
    isLoading.value = true;

    fetchData();
}

onMounted(() => {
    isLoading.value = true;
    fetchData();
});

watchPostEffect(() => {
    isLoading.value = true;
    fetchData();
});

</script >

<template lang="">
    <aside>
        <table class="w-full text-left border border-black ">
            <thead class="bg-gradient-to-r from-cyan-500 to-blue-600 text-white ">
                <tr>
                    <th class="py-1 px-2" v-for="(item, index) of props.definition" :key="index">
                        {{ item.label }}
                    </th>
                </tr>
            </thead>
            <tbody  v-if="data.length && !isLoading">
                <tr v-for="row of data" class="border-b border-black cursor-pointer hover:bg-gray-200">
                    <td class="py-1 px-" @click="emitClick(row)" v-for="{key, transform} of props.definition">
                        <p class="">
                            {{ transform ? transform(row[key]) : row[key] }}
                        </p>
                    </td>
                </tr>
            </tbody>
        </table>
        <div v-if="isLoading" class="flex justify-center pt-6">
            <Spinner />
        </div>

        <div v-if="!isLoading && data.length" class="flex gap-2 mt-2">
            <h3 class="font-bold">Current Page <span class="font-normal">{{ currentPage }}</span></h3>
            <div class="flex gap-2">
                <button @click="prevPage" class="bg-blue-500 text-white font-bold px-2 py-1 rounded-lg">Prev</button>
                <button @click="nextPage" class="bg-blue-500 text-white font-bold px-2 py-1 rounded-lg">Next</button>
            </div>
            <select class="font-bold" v-model="pageSize">
                <option v-for="option of props.pageSizeOptions" :value="option">{{option}}</option>
            </select>
        </div>
        <div v-if="!data.length && !isLoading" class="flex justify-center  pt-4">
            <h1 class="font-bold text-2xl">No Data</h1>
        </div>
    </aside>
</template>