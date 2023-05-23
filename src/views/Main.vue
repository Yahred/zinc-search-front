<script setup lang="ts">
import { ref, watchPostEffect } from 'vue'


import type { TableHeader } from '@/interfaces';

import { SEARCH } from '../config/endpoints'
import Visualizer from '../components/Visualizer.vue';
import Table from '../components/Table.vue'

const tableDefinition: TableHeader[] = [
    {
        key: 'subject',
        label: 'Subject'
    },
    {
        key: 'from',
        label: 'From',
    },
    {
        key: 'to',
        label: 'To',
    },
]


const search = ref('');
const selectedMail = ref('');
const tableUri = ref(SEARCH);

function tableClick(item: any) {
    selectedMail.value = item.body;
}

watchPostEffect(() => {
    console.log(search.value);
});


</script>
<template>
    <main class="w-full flex flex-col gap-4 p-4">
        <input v-model="search"
            class="w-full drop-shadow-xl border-2 h-12 p-4 rounded-lg focus:ring-blue-500 focus:border-blue-500"
            placeholder="Search" type="text">

        <section class="grid xl:grid-cols-2 grid-cols-1 gap-2">
            <Table @click="tableClick" :definition="tableDefinition" :search="search" :uri="tableUri" />
            <Visualizer :email="selectedMail"/>
        </section>
    </main>
</template>