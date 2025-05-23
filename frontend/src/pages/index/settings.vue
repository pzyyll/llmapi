<script setup lang="ts">
const { locale, locales, setLocale } = useI18n();
</script>

<template>
	<MainLayout :title="$t('settings')">
		<MainHeader :title="$t('settings')" />
		<MainContent>
			<div class="flex flex-col gap-4 p-4">
				<div name="language" class="flex flex-col gap-2">
					<div class="flex flex-col gap-2">
						<fieldset class="d-fieldset max-w-60">
							<legend class="d-fieldset-legend text-sm">{{ $t('language') }}</legend>
							<select class="d-select d-select-sm" v-model="locale" @change="setLocale(locale)">
								<option v-for="item in locales" :value="item.code">
									{{ item.name }}
								</option>
							</select>
						</fieldset>
					</div>
				</div>
			</div>
		</MainContent>
	</MainLayout>
</template>
