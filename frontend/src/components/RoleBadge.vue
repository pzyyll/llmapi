<script setup lang="ts">
const props = defineProps({
	role: { type: String, default: '' }
});
const roleColorMap = {
	[Role.Super]: 'd-badge-error',
	[Role.Admin]: 'd-badge-warning',
	[Role.User]: 'd-badge-success'
};
</script>

<template>
	<span :class="roleColorMap[props.role as Role]" class="d-badge d-badge-sm whitespace-nowrap">
		{{ $t(RoleNames[props.role as Role]) }}
	</span>
</template>
