<script>
	import { Tabs, TabItem } from 'flowbite-svelte';
	import GeneralTab from './GeneralTab.svelte';

	export let data;
</script>

<Tabs tabStyle="underline" contentClass="p-4 mt-4">
	<TabItem open title="General">
		<GeneralTab form={data.form} questions={data.questions} />
	</TabItem>
	<TabItem title="View" disabled></TabItem>
	<TabItem title="Responses" disabled></TabItem>
	<TabItem title="History" disabled></TabItem>
</Tabs>
