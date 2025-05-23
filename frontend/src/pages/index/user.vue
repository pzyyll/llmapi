<script setup lang="ts">
const users = ref<Dto.Users>();
const api = useAPI();
const { isSelf } = useAuthStore();

async function handleDelete(userId: string) {
	console.log('Delete user with ID:', userId);
	// Remove the user from the local state
	try {
		await api.delete(RequestPath.DeleteUser, {
			params: {
				user_id: userId
			}
		});
		users.value!.users = users.value!.users.filter((user) => user.user_id !== userId);
	} catch (error) {
		console.error('Delete User Error', error);
	}
}

onMounted(async () => {
	// get users from API
	try {
		const { data } = await api.get<Dto.Users>(RequestPath.Users);
		console.log('data: ', data);
		users.value = data;
	} catch (error) {
		console.error('Get Users Error', error);
	}
});
</script>

<template>
	<MainLayout>
		<MainHeader :title="$t('user')" />
		<MainContent>
			<div class="container mx-auto flex max-w-5xl flex-col">
				<div class="rounded-box border-base-content/10 bg-base-100 overflow-x-auto border">
					<table class="d-table w-full">
						<!-- head -->
						<colgroup>
							<col class="w-2/5" />
							<col class="w-1/5" />
							<col class="w-1/5" />
						</colgroup>
						<tbody>
							<tr
								v-for="user in users?.users"
								:key="user.user_id"
								class="border-base-content/10 border-b"
							>
								<td>
									<div class="flex items-center gap-2">
										<AvatarIcon :name="user.username" />
										<div class="flex items-center gap-1">
											<p>
												{{ user.username }}
											</p>
											<div
												class="d-badge d-badge-info bg-info/30 d-badge-xs border-0"
												v-if="isSelf(user.user_id)"
											>
												{{ $t('you') }}
											</div>
										</div>
									</div>
								</td>
								<td>
									<RoleBadge :role="user.role" />
								</td>
								<td class="flex items-center justify-end">
									<ConfirmDialog @confirm="handleDelete(user.user_id)" :title="$t('delete_user')">
										<button
											class="d-btn d-btn-ghost d-btn-xs hover:bg-error/20 active:bg-error/20 focus:outline-error hover:border-transparent focus:border-transparent focus:bg-transparent active:border-transparent"
										>
											<MaterialSymbolsDeleteOutlineRounded class="text-error size-4" />
										</button>
										<template #message>
											<i18n-t
												keypath="delete_user_message"
												tag="p"
												class="text-base-content/80 my-1 text-sm"
											>
												<template #username>
													<span class="font-bold italic text-black">{{ user.username }}</span>
												</template>
											</i18n-t>
										</template>
									</ConfirmDialog>
								</td>
							</tr>
						</tbody>
					</table>
				</div>
			</div>
		</MainContent>
	</MainLayout>
</template>
