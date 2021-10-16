<script>
	import { Accordion, AccordionItem } from 'sveltestrap';
</script>

<div class="container">
	<h1 class="mt-3">Home</h1>

	<Accordion>
		<AccordionItem active header="ZIM">Explain ZIM Page</AccordionItem>
		<AccordionItem active header="Error">Explain Error Page</AccordionItem>
        <AccordionItem active header="Sensor">Explain Sensor</AccordionItem>
        <AccordionItem active header="Form">Explain Form</AccordionItem>
        <AccordionItem active header="Database">Explain Database</AccordionItem>
        <AccordionItem active header="Ajax">Explain Ajax</AccordionItem>
	</Accordion>
</div>
