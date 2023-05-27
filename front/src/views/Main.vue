<script setup lang="ts">
import { ref, watchPostEffect } from 'vue'


import type { TableHeader } from '@/interfaces';

import Visualizer from '../components/Visualizer.vue';
import Table from '../components/Table.vue'
import Footer from '@/layout/Footer.vue';

import { SEARCH } from '../config/endpoints'

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

        <section class="grid xl:grid-cols-2 grid-cols-1 gap-2 max-h-[80vh]">
            <Table @click="tableClick" :definition="tableDefinition" :search="search" :uri="tableUri" />
            <div class="max-h-[35vh] lg:max-h-[70vh] overflow-y-auto">
                <Visualizer :keywords="search.split(' ')" :email="selectedMail"/>
            </div>
        </section>
        
    </main>
    <Footer />
</template>