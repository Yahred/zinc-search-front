<script setup lang="ts">
import { ref, onMounted, watchPostEffect, withDefaults } from 'vue'

import { type TableHeader } from '../interfaces'
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
    console.log(response);

    data.value = response;
}


function nextPage() {
    if (data.value.length < pageSize.value) return;

    currentPage.value++;
    fetchData();
}

function prevPage(){
    if (currentPage.value === 1) return;

    currentPage.value--;
    fetchData();
}

onMounted(() => {
    fetchData();
});

watchPostEffect(() => {
    fetchData();
});

</script >

<template lang="">
    <aside>
        <table class="w-full text-left border border-black ">
            <thead class="bg-blue-500 text-white ">
                <tr>
                    <th class="py-1 px-2" v-for="(item, index) of props.definition" :key="index">
                        {{ item.label }}
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row of data" class="border-b border-black cursor-pointer hover:bg-gray-200">
                    <td class="py-1 px-2" @click="emitClick(row)" v-for="{key, transform} of props.definition">
                        {{ transform ? transform(row[key]) : row[key] }}
                    </td>
                </tr>
            </tbody>
        </table>
        <div class="flex gap-2 mt-2">
            <h3>Current Page {{currentPage}}</h3>
            <div class="flex gap-2">
                <button @click="prevPage" class="bg-blue-500 text-white px-2 py-1 rounded-lg">Prev</button>
                <button @click="nextPage" class="bg-blue-500 text-white px-2 py-1 rounded-lg">Next</button>
            </div>
            <select v-model="pageSize">
                <option v-for="option of props.pageSizeOptions" :value="option">{{option}}</option>
            </select>
        </div>
    </aside>
</template>