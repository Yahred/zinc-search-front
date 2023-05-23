<script setup lang="ts">
import {
    watchEffect
    , ref, type VNodeRef
} from 'vue'

interface VisualizerProps {
    email: string
}

const props = defineProps<VisualizerProps>();

const pRef = ref<VNodeRef | null>(null);

watchEffect(() => {
    console.log(pRef.value);
    if (pRef.value) {
        pRef.value.innerHTML = props.email;
    }

    console.log(props.email);
});


</script>
<template>
    <p :ref="pRef"></p>
</template>
