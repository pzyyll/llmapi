<script setup lang="ts">
const { getUser } = useAuthStore();

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
				<AvatarIcon :name="getUser?.username" size="sm" />
			</DropdownMenuTrigger>

			<DropdownMenuPortal>
				<DropdownMenuContent
					class="d-menu bg-base-100 shadow-xl/30 w-32 overflow-clip rounded-md px-0"
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
