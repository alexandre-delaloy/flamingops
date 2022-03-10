import { useState, useEffect, useCallback, FormEvent, useRef } from 'react';
import type { ChangeEvent, MouseEvent } from 'react';
import type { NextPage } from 'next';
import Head from 'next/head';
import {
  Box,
  Container,
  Grid,
  Typography
} from '@mui/material';
// import { customerApi } from '../../../api/customer-api';
import { AuthGuard } from '../../../components/authentication/auth-guard';
import { DashboardLayout } from '../../../components/dashboard/dashboard-layout';
// import { OverviewLatestTransactions } from '../../../components/dashboard/overview/overview-latest-transactions';
import { OverviewLatestTransactionsScaleway } from '../../../components/dashboard/overview/overview-latest-transactions-scaleway';
import { useMounted } from '../../../hooks/use-mounted';
import { gtm } from '../../../lib/gtm';
// import type { Customer } from '../../../types/customer';

interface Filters {
  query?: string;
  hasAcceptedMarketing?: boolean;
  isProspect?: boolean;
  isReturning?: boolean;
}


const applyFilters = (
  customers: Customer[],
  filters: Filters
): Customer[] => customers.filter((customer) => {
  if (filters.query) {
    let queryMatched = false;
    const properties = ['email', 'name'];

    properties.forEach((property) => {
      if (customer[property].toLowerCase().includes(filters.query.toLowerCase())) {
        queryMatched = true;
      }
    });

    if (!queryMatched) {
      return false;
    }
  }

  if (filters.hasAcceptedMarketing && !customer.hasAcceptedMarketing) {
    return false;
  }

  if (filters.isProspect && !customer.isProspect) {
    return false;
  }

  if (filters.isReturning && !customer.isReturning) {
    return false;
  }

  return true;
});

const descendingComparator = (
  a: Customer,
  b: Customer,
  orderBy: string
): number => {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }

  if (b[orderBy] > a[orderBy]) {
    return 1;
  }

  return 0;
};

const getComparator = (order: 'asc' | 'desc', orderBy: string) => (
  order === 'desc'
    ? (a: Customer, b: Customer) => descendingComparator(a, b, orderBy)
    : (a: Customer, b: Customer) => -descendingComparator(a, b, orderBy)
);

const applyPagination = (
  customers: Customer[],
  page: number,
  rowsPerPage: number
): Customer[] => customers.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage);

const CustomerList: NextPage = () => {
  const isMounted = useMounted();
  const queryRef = useRef<HTMLInputElement | null>(null);
  const [customers, setCustomers] = useState<Customer[]>([]);
  const [currentTab, setCurrentTab] = useState<string>('all');
  const [page, setPage] = useState<number>(0);
  const [rowsPerPage, setRowsPerPage] = useState<number>(10);
  // const [sort, setSort] = useState<Sort>(sortOptions[0].value);
  const [filters, setFilters] = useState<Filters>({
    query: '',
    hasAcceptedMarketing: null,
    isProspect: null,
    isReturning: null
  });

  useEffect(() => {
    gtm.push({ event: 'page_view' });
  }, []);

  const getCustomers = useCallback(async () => {
    try {
      const data = await customerApi.getCustomers();

      if (isMounted()) {
        setCustomers(data);
      }
    } catch (err) {
      console.error(err);
    }
  }, [isMounted]);

  useEffect(
    () => {
      getCustomers();
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const handleTabsChange = (event: ChangeEvent<{}>, value: string): void => {
    const updatedFilters = {
      ...filters,
      hasAcceptedMarketing: null,
      isProspect: null,
      isReturning: null
    };

    if (value !== 'all') {
      updatedFilters[value] = true;
    }

    setFilters(updatedFilters);
    setCurrentTab(value);
  };

  const handleQueryChange = (event: FormEvent<HTMLFormElement>): void => {
    event.preventDefault();
    setFilters((prevState) => ({
      ...prevState,
      query: queryRef.current?.value
    }));
  };

  const handlePageChange = (event: MouseEvent<HTMLButtonElement> | null, newPage: number): void => {
    setPage(newPage);
  };

  const handleRowsPerPageChange = (event: ChangeEvent<HTMLInputElement>): void => {
    setRowsPerPage(parseInt(event.target.value, 10));
  };

  return (
    <>
      <Head>
        <title>
          Dashboard: Scaleway List
        </title>
      </Head>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          py: 8
        }}
      >
        <Container maxWidth="xl">
          <Box sx={{ mb: 4 }}>
              <Grid
                container
                justifyContent="space-between"
                spacing={3}
              >
                <Grid item>
                  <Typography variant="h4">
                    Scaleway
                  </Typography>
                </Grid>
              </Grid>
            </Box>
          <OverviewLatestTransactionsScaleway />
        </Container>
      </Box>
    </>
  );
};

CustomerList.getLayout = (page) => (
  <AuthGuard>
    <DashboardLayout>
      {page}
    </DashboardLayout>
  </AuthGuard>
);

export default CustomerList;
