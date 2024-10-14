<script lang="ts">
	import type { Form } from '$lib/proto/form/v1/forms_pb.js';

	// import type { Form } from '$lib/proto/form/v1/forms';
	import { Card, Badge, Indicator, Button } from 'flowbite-svelte';

	export let form: Form;
</script>

<Card href="view/{form.baseId}" class="min-w-[300px]">
	<div class="pepetest mb-4 flex justify-between">
		<h2 class="text-lg font-bold text-gray-900 dark:text-gray-100">{form.title}</h2>
		<Badge color="green" rounded class="px-2.5 py-0.5">
			<!-- <Indicator color="green" size="xs" class="me-1" /> -->
			Active
		</Badge>
	</div>

	<p class="mb-4 text-sm">{form.description}</p>

	<!-- <Button>Manage</Button> -->
</Card>
