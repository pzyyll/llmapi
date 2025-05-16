<script setup lang="ts">
import gsap from 'gsap';

const list = [
	{ msg: 'Bruce Lee' },
	{ msg: 'Jackie Chan' },
	{ msg: 'Chuck Norris' },
	{ msg: 'Jet Li' },
	{ msg: 'Kung Fury' }
];

const query = ref('');

const computedList = computed(() => {
	return list.filter((item) => {
		return !item.msg.toLowerCase().includes(query.value);
	});
});

function onEnter(el: Element, done: () => void) {
	gsap.from(el, {
		opacity: 0,
		height: 0,
		delay: el.dataset.index * 0.6,
		onComplete: done
	});
}

function onLeave(el: Element, done: () => void) {
	gsap.to(el, {
		opacity: 0,
		height: 0,
		delay: el.dataset.index * 0.6,
		onComplete: done
	});
}
</script>

<template>
	<input v-model="query" />
	<TransitionGroup
		tag="ul"
		:css="false"
		@enter="onEnter"
		@leave="onLeave"
	>
		<li v-for="(item, index) in computedList" :key="item.msg" :data-index="index">
			{{ item.msg }}
		</li>
	</TransitionGroup>
</template>
