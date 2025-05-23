<template>
	<div class="h-full overflow-auto">
		<div>
			<slot name="header">
				<!-- <div class="pt-4"></div> -->
			</slot>
			<div class="px-5 py-4">
				<slot />
			</div>
		</div>
	</div>
</template>
