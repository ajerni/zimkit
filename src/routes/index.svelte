<script>
    import { Accordion, AccordionItem, Collapse, Button, Card } from 'sveltestrap';
    let isOpen = false;
</script>

<div class="container">

<h1 class="mt-3">Home</h1>

  <Accordion>
    <AccordionItem active header="Home">Fallbrook</AccordionItem>
    <AccordionItem header="School">
      <a href="#home">Buena Vista Elementary</a>
    </AccordionItem>
    <AccordionItem header="Library">UCSB Library</AccordionItem>
  </Accordion>

  <hr>

  <Button color="primary" on:click={() => (isOpen = !isOpen)} class="mb-3">
    Toggle
  </Button>
  <Collapse {isOpen}>
    <Card body>
      Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry
      richardson ad squid. Nihil anim keffiyeh helvetica, craft beer labore wes
      anderson cred nesciunt sapiente ea proident.
    </Card>
  </Collapse>

</div>
