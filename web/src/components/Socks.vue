<script setup lang="ts">
import { onMounted, ref, computed } from "vue";
import { createOrder, type OrderSummary } from "@/services/orderService";
import { getProduct, type Product } from "@/services/productService";

const product = ref<Product | null>(null);
const quantity = ref<number>(1);
const orderSummary = ref<OrderSummary | null>(null);
const orderForm = ref<HTMLFormElement | null>(null);


function validateQuantity(event: Event) {
  const input = event.target as HTMLInputElement;

  input.setCustomValidity(
    input.validity.rangeOverflow ? "Come on, be cool!" : "",
  );
}

async function loadProduct(id: number) {
  try {
    product.value = await getProduct(id);
  } catch (err) {
    // TODO: Error handling here
    console.error(err)
  }
}

async function submitOrder() {
  if (!orderForm.value?.reportValidity()) return;
  if (!product.value) return;

  try {
    orderSummary.value = await createOrder({
      productId: product.value.id,
      quantity: quantity.value,
    });
  } catch (err) {
    // TODO: Error handling here
    console.error(err)
  }
}

const socksYouDidntWant = computed(() => {
  if (orderSummary.value !== null)
    return orderSummary.value.actualQuantity - quantity.value

  return null
})

onMounted(() => {
  loadProduct(4);
});
</script>

<template>
  <div v-if="product">
    <p>We only sell: {{product?.name}}</p>
    <p>In the following pack sizes:</p>
    <ul v-if="product?.packSizes">
      <li v-for="pack in product.packSizes" :key="pack.size">
        {{pack.size}} @ £{{pack.cost}}
      </li>
    </ul>
    <br>
    <br>
    <p>Enter your desired quantity below, and confirm your order</p>
    <form @submit.prevent="submitOrder" ref="orderForm">
      <label>Quantity:</label>
      <input ref="quantityInput" required v-model="quantity" type="number" min="1" max="99999" @input="validateQuantity">
      <br>
      <br>
      <button class="button" type="submit" @click="submitOrder">Submit</button>
      <br>
      <br>
    </form>

    <div class="orderSummary" v-if="orderSummary">
      <p>Congrats, you've got more socks coming than you possibly know what to do with!</p>
      <p>Order Summary</p>
      <p>OrderId: {{orderSummary.orderId}}</p>
      <p>ProductId: {{orderSummary.productId}}</p>
      <p>Quantity you wanted: {{orderSummary.desiredQuantity}}</p>
      <p>Quantity you're paying for: {{orderSummary.actualQuantity}}</p>
      <p>Socks you didn't want, but are getting: {{socksYouDidntWant}}</p>
      <p>Cost: £{{orderSummary.cost}}</p>
      <p>How we determined the cheapest way to allocate the packs of socks:</p>
      <ul>
        <li v-for="allocation in orderSummary.packAllocation">
          {{allocation.packSize}} socks in pack. Quantity: {{allocation.quantity}} @ £{{allocation.packCost}}
        </li>
      </ul>
    </div>
  </div>

</template>

<style scoped>
.orderSummary {
  /* This can later become a themed CSS custom property. */
  background-color: rgba(221, 255, 221, 0.5);
}
</style>
