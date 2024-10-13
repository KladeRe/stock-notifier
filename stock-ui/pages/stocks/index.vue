<script setup lang="ts">
interface StockBody {
    Id: Number
    Symbol: string,
    Value: String,
    Buy: Boolean,
    Notification: string
}
const { status, data: stocks} = await useLazyFetch<StockBody[]>("http://localhost:5050/configs")
watch(stocks, (newStocks) => {})
</script>

<template>
  <div v-if="status === 'pending'">
    Loading ...
  </div>
  <div v-else>
    <div>Stocks loaded</div>
    <div v-for="stock in stocks">
    <p>{{ stock.Symbol }}</p>
      
      <ul>
        <li>
            <p>{{ stock.Value }}</p>
        </li>
      </ul>
    </div>
  </div>
</template>