<script setup lang="ts">
const handleExit = async () => {
	try {
		await useAPI().post(RequestPath.Logout);
		const authStore = useAuthStore();
		authStore.clearAuth();
		navigateTo('/login');
	} catch (error) {
		console.error('Logout error:', error);
	}
};
</script>

<template>
	<div>
		<DropdownMenuRoot>
			<DropdownMenuTrigger asChild>
				<div class="d-avatar d-avatar-placeholder">
					<button class="d-btn d-btn-info d-btn-sm d-btn-circle">
						<span class="text-base-100 text-xl">A</span>
					</button>
				</div>
			</DropdownMenuTrigger>

			<DropdownMenuPortal>
				<DropdownMenuContent
					class="d-menu bg-base-100 rounded-box shadow-xl/30 w-32 px-0"
					align="end"
					:alignOffset="-2"
					side="bottom"
					:sideOffset="2"
				>
					<DropdownMenuItem
						class="d-btn d-btn-ghost d-btn-xs flex items-center justify-start rounded-none"
						@click="handleExit"
					>
						<MaterialSymbolsLogoutRounded />
						<span class="text-xs font-normal">{{ $t('logout') }}</span>
					</DropdownMenuItem>
					<DropdownMenuArrow class="fill-base-100" />
				</DropdownMenuContent>
			</DropdownMenuPortal>
		</DropdownMenuRoot>
	</div>
</template>
