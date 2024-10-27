import { Text, StyleSheet, Platform } from 'react-native';
import { ThemedView } from '@/components/ThemedView';
import { useEffect, useState } from 'react';

interface StockType {
  id: string,
  symbol: string,
  value: number,
  buy: boolean,
  notification: string
}

interface APIFetchState {
  data: StockType[],
  loading: boolean,
  error: string | null
}

export default function StocksScreen() {
  const [fetchState, setFetchState] = useState<APIFetchState>({
    data: [],
    loading: true,
    error: null
  })
  const fetchStocks = async () => {
    try {
      const response = await fetch(
        `http://localhost:5050/configs`
      );
      if (!response.ok) {
        throw new Error(`HTTP error: Status ${response.status}`);
      }
      let postsData: StockType[] = await response.json();
      setFetchState({
        data: postsData,
        loading: false,
        error: null
      });
    } catch (err) {
      setFetchState({
        data: [],
        loading: false,
        error: (err as Error).message
      })
    } 
  }

  useEffect(() => {
    fetchStocks()
  }, [])
  return (
    <div>
      {fetchState.loading && <p>Loading data</p>}
      {fetchState.error && <p>Error: {fetchState.error}</p>}
      { fetchState.data.length > 0 ? (
      <ul>
        {fetchState.data.map((stock) => (
          <li key={stock.id}>
            <h3>{stock.symbol}</h3>
            <p>{stock.value}</p>
            <p>{String(stock.buy)}</p>
            <p>{stock.notification}</p>
          </li>
        ))}
      </ul>
      ) : 
      <p>No stocks found</p>
    }
    </div>
  )
    
    
}

const styles = StyleSheet.create({
    titleContainer: {
      flexDirection: 'row',
      alignItems: 'center',
      gap: 8,
    },
    stepContainer: {
      gap: 8,
      marginBottom: 8,
    },
    reactLogo: {
      height: 178,
      width: 290,
      bottom: 0,
      left: 0,
      position: 'absolute',
    },
  });