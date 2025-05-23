<script setup lang="ts"></script>

<template>
	<MainContent>
		<div class="d-tabs d-tabs-border">
			<input type="radio" name="my_tabs_2" class="d-tab" aria-label="Tab 1" />
			<div class="d-tab-content border-base-300 bg-base-100 p-10">Tab content 1</div>

			<input type="radio" name="my_tabs_2" class="d-tab" aria-label="Tab 2" checked />
			<div class="d-tab-content border-base-300 bg-base-100 p-10">Tab content 2</div>

			<input type="radio" name="my_tabs_2" class="d-tab" aria-label="Tab 3" />
			<div class="d-tab-content border-base-300 bg-base-100 p-10">Tab content 3</div>
		</div>
	</MainContent>
</template>
