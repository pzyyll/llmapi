<script setup lang="ts">
const users = ref<Dto.Users>();

function handleDelete(userId: number) {
	console.log('Delete user with ID:', userId);
	// Remove the user from the local state
	users.value!.users = users.value!.users.filter((user) => user.user_id !== userId);
	// Optionally, you can show a success message or refresh the user list
}

onMounted(async () => {
	// get users from API
	try {
		const { data } = await useAPI().get<Dto.Users>(RequestPath.Users);
		console.log('data: ', data);
		users.value = data;
	} catch (error) {
		console.error('Get Users Error', error);
	}
});
</script>

<template>
	<DashboardMainLayout :title="$t('user')">
		<div class="overflow-x-auto">
			<table class="d-table d-table-zebra w-full">
				<!-- head -->
				<thead>
					<tr>
						<th>{{ $t('username') }}</th>
						<th>{{ $t('role') }}</th>
						<th>{{ $t('created_at') }}</th>
						<th>{{ $t('edit') }}</th>
					</tr>
				</thead>
				<tbody>
					<!-- row 1 -->
					<tr v-for="user in users?.users" :key="user.user_id">
						<td>{{ user.username }}</td>
						<td>
							<RoleBadge :role="user.role" />
						</td>
						<td>
							{{ new Date(user.created_at).toLocaleDateString() }}
						</td>
						<td>
							<button
								class="d-btn d-btn-ghost d-btn-xs d-btn-error"
								@click="handleDelete(user.user_id)"
							>
								<MaterialSymbolsDeleteOutlineRounded class="size-4" />
							</button>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</DashboardMainLayout>
</template>
