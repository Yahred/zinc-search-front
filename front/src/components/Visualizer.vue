<script setup lang="ts">
import {
    watchEffect
    , ref, type VNodeRef
} from 'vue'

interface VisualizerProps {
    email: string,
    keywords?: string[],
}

const props = defineProps<VisualizerProps>();

const pRef = ref<VNodeRef | null>(null);

function highLightKeywords() {
    if (!props.keywords) return;

    const regex = new RegExp(props.keywords.join('|'), 'gi');
    const text = pRef.value?.innerHTML || '';

    pRef.value!.innerHTML = text.replace(regex, (match: string) => `<span class="bg-yellow-200">${match}</span>`);
}

watchEffect(() => {
    if (!pRef.value) return;
    
    pRef.value.innerHTML = props.email;
    highLightKeywords();

});


</script>
<template>
    <p ref="pRef"></p>
</template>
