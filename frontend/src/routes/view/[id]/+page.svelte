<script>
	import { Tabs, TabItem } from 'flowbite-svelte';
	import GeneralTab from './GeneralTab.svelte';
	import ViewTab from './ViewTab.svelte';

	export let data;
</script>

<Tabs tabStyle="underline" contentClass="p-4 mt-4 min-w-[50%]">
	<TabItem open title="General">
		<GeneralTab form={data.form} questions={data.questions} />
	</TabItem>
	<TabItem title="View">
		<ViewTab id={data.form.baseId} />
	</TabItem>
	<TabItem title="Responses" disabled></TabItem>
	<TabItem title="History" disabled></TabItem>
</Tabs>
