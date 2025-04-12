<template>
	<div class="bg-base-100 rounded-md h-full">
		<h1><IcRoundHome />Welcome to the homepage</h1>
		<AppAlert> This is an auto-imported component </AppAlert>
		<nav>
			<ul>
				<li><NuxtLink to="/about">About</NuxtLink></li>
				<li><NuxtLink to="/demo/demopage">Post 1</NuxtLink></li>
				<li><NuxtLink to="/demo/demopage2">Post 2</NuxtLink></li>
			</ul>
		</nav>
	</div>
</template>
