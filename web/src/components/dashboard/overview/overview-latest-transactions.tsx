import type { FC } from 'react';
import { format, subDays } from 'date-fns';
import numeral from 'numeral';
import {
  Box,
  Card,
  CardHeader,
  Table,
  TableBody,
  TableHead,
  TableCell,
  TableRow,
  Typography
} from '@mui/material';
import { Scrollbar } from '../../scrollbar';
import { SeverityPill } from '../../severity-pill';

interface Transaction {
  id: string;
  name: string;
  type: string;
  status: string;
  creationDate: string;
  modificationDate: string;
  cpuLoad: string;
  ramUsage: string;
}

const transactions: Transaction[] = [
  {
    id: 'd46800328cd510a668253b45',
    name: 'AWS',
    type: 'IoT',
    status: 'running',
    creationDate: '3 months ago',
    modificationDate: '3 months ago',
    cpuLoad: '15%',
    ramUsage: 'WARNING 80%'
  },
  {
    id: 'd46800328cderz0a668253b45',
    name: 'Scaleway',
    type: 'EC2',
    status: 'running',
    creationDate: '2 days ago',
    modificationDate: '2 days ago',
    cpuLoad: '15%',
    ramUsage: '20%'
  },
  {
    id: 'd4682328cd510a668253b45',
    name: 'AWS',
    type: 'IoT',
    status: 'running',
    creationDate: '1 month ago',
    modificationDate: 'two weeks ago',
    cpuLoad: 'WARNING 80%',
    ramUsage: '20%'
  },
  {
    id: 'd46800328cd510a6682ghb45',
    name: 'Azure',
    type: 'IoT',
    status: 'running',
    creationDate: '2 weeks ago',
    modificationDate: '1 week ago',
    cpuLoad: '15%',
    ramUsage: '20%'
  },
  {
    id: 'd46800328cd510a668253b42',
    name: 'AWS',
    type: 'S3 Bucket',
    status: 'running',
    creationDate: '5 days ago',
    modificationDate: '5 days ago',
    cpuLoad: '15%',
    ramUsage: '20%'
  },
];

export const OverviewLatestTransactions: FC = (props) => (
  <Card {...props}>
    <CardHeader title="Instances" />
    <Scrollbar>
      <Table sx={{ minWidth: 600 }}>
        <TableHead>
          <TableRow >
            <TableCell>
              Name
            </TableCell>
            <TableCell>
              Type
            </TableCell>
            <TableCell>
              Status
            </TableCell>
            <TableCell>
              Creation date
            </TableCell>
            <TableCell>
              Modification date
            </TableCell>
            <TableCell>
              CPU load
            </TableCell>
            <TableCell>
              RAM usage
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {transactions.map((transaction) => (
            <TableRow
              key={transaction.id}
              sx={{
                '&:last-child td': {
                  border: 0
                }
              }}
            >
              <TableCell width={300}>
                <Box
                  sx={{
                    p: 1,
                    backgroundColor: (theme) => theme.palette.mode === 'dark'
                      ? 'neutral.800'
                      : 'neutral.200',
                    borderRadius: 2,
                    maxWidth: 'fit-content'
                  }}
                >
                  <Typography
                    align="center"
                    color="textSecondary"
                    variant="subtitle2"
                  >
                    {transaction.name}
                  </Typography>
                </Box>
              </TableCell>
              <TableCell>
                <div>
                  <Typography >
                    {transaction.type}
                  </Typography>
                </div>
              </TableCell>
              <TableCell>
                <div>
                  <SeverityPill>
                    {transaction.status}
                  </SeverityPill>
                </div>
              </TableCell>
              <TableCell width={180}>
                <Typography
                  variant="subtitle2"
                >
                  {transaction.creationDate}
                </Typography>
              </TableCell>
              <TableCell width={180}>
                <Typography
                  variant="subtitle2"
                >
                  {transaction.modificationDate}
                </Typography>
              </TableCell>
              <TableCell width={180}>
                <Typography
                  variant="subtitle2"
                >
                  {transaction.cpuLoad}
                </Typography>
              </TableCell>
              <TableCell width={180}>
                <Typography
                  variant="subtitle2"
                >
                  {transaction.ramUsage}
                </Typography>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Scrollbar>
  </Card>
);
