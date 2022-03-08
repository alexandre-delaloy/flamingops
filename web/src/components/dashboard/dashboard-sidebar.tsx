import { ReactNode, useEffect, useMemo, useRef, useState } from 'react';
import type { FC } from 'react';
import NextLink from 'next/link';
import { useRouter } from 'next/router';
import PropTypes from 'prop-types';
import { useTranslation } from 'react-i18next';
import type { TFunction } from 'react-i18next';
import { Box, Divider, Drawer, Typography, useMediaQuery } from '@mui/material';
import type { Theme } from '@mui/material';
import { Calendar as CalendarIcon } from '../../icons/calendar';
import { Cash as CashIcon } from '../../icons/cash';
import { ChartBar as ChartBarIcon } from '../../icons/chart-bar';
import { ChartPie as ChartPieIcon } from '../../icons/chart-pie';
import { ChatAlt2 as ChatAlt2Icon } from '../../icons/chat-alt2';
import { ClipboardList as ClipboardListIcon } from '../../icons/clipboard-list';
import { CreditCard as CreditCardIcon } from '../../icons/credit-card';
import { Home as HomeIcon } from '../../icons/home';
import { LockClosed as LockClosedIcon } from '../../icons/lock-closed';
import { Mail as MailIcon } from '../../icons/mail';
import { MailOpen as MailOpenIcon } from '../../icons/mail-open';
import { Newspaper as NewspaperIcon } from '../../icons/newspaper';
import { OfficeBuilding as OfficeBuildingIcon } from '../../icons/office-building';
import { ReceiptTax as ReceiptTaxIcon } from '../../icons/receipt-tax';
import { Selector as SelectorIcon } from '../../icons/selector';
import { Share as ShareIcon } from '../../icons/share';
import { ShoppingBag as ShoppingBagIcon } from '../../icons/shopping-bag';
import { ShoppingCart as ShoppingCartIcon } from '../../icons/shopping-cart';
import { Truck as TruckIcon } from '../../icons/truck';
import { UserCircle as UserCircleIcon } from '../../icons/user-circle';
import { Users as UsersIcon } from '../../icons/users';
import { XCircle as XCircleIcon } from '../../icons/x-circle';
import { Logo } from '../logo';
import { Scrollbar } from '../scrollbar';
import { DashboardSidebarSection } from './dashboard-sidebar-section';
import { OrganizationPopover } from './organization-popover';

interface DashboardSidebarProps {
  onClose: () => void;
  open: boolean;
}

interface Item {
  title: string;
  children?: Item[];
  chip?: ReactNode;
  icon?: ReactNode;
  path?: string;
}

interface Section {
  title: string;
  items: Item[];
}

const getSections = (t: TFunction): Section[] => [
  {
    title: t('General'),
    items: [
      {
        title: t('Overview'),
        path: '/dashboard',
        icon: <HomeIcon fontSize="small" />
      },
      {
        title: t('Analytics'),
        path: '/dashboard/analytics',
        icon: <ChartBarIcon fontSize="small" />
      },
      {
        title: t('Account'),
        path: '/dashboard/account',
        icon: <UserCircleIcon fontSize="small" />
      }
    ]
  },
  {
    title: t('Management'),
    items: [
      {
        title: t('Scaleway settings'),
        path: '/dashboard/scaleway',
        icon: <UsersIcon fontSize="small" />,
        children: [
          {
            title: t('List'),
            path: '/dashboard/scaleway'
          },
          {
            title: t('Details'),
            path: '/dashboard/scaleway/1'
          },
          {
            title: t('Edit'),
            path: '/dashboard/scaleway/1/edit'
          }
        ]
      },
      {
        title: t('AWS settings'),
        path: '/dashboard/products',
        icon: <UsersIcon fontSize="small" />,
        children: [
          {
            title: t('List'),
            path: '/dashboard/AWS'
          },
          {
            title: t('Details'),
            path: '/dashboard/AWS/1'
          },
          {
            title: t('Edit'),
            path: '/dashboard/AWS/1/edit'
          }
        ]
      },
      {
        title: t('Azure settings'),
        icon: <UsersIcon fontSize="small" />,
        path: '/dashboard/orders',
        children: [
          {
            title: t('List'),
            path: '/dashboard/azure'
          },
          {
            title: t('Details'),
            path: '/dashboard/azure/1'
          },
          {
            title: t('Edit'),
            path: '/dashboard/azure/1/edit'
          }
        ]
      },
      {
        title: t('GCP settings'),
        path: '/dashboard/invoices',
        icon: <ReceiptTaxIcon fontSize="small" />,
      }
    ]
  },
];

export const DashboardSidebar: FC<DashboardSidebarProps> = (props) => {
  const { onClose, open } = props;
  const router = useRouter();
  const { t } = useTranslation();
  const lgUp = useMediaQuery(
    (theme: Theme) => theme.breakpoints.up('lg'),
    {
      noSsr: true
    }
  );
  const sections = useMemo(() => getSections(t), [t]);

  const handlePathChange = () => {
    if (!router.isReady) {
      return;
    }

    if (open) {
      onClose?.();
    }
  };

  useEffect(
    handlePathChange,
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [router.isReady, router.asPath]
  );

  const content = (
    <>
      <Scrollbar
        sx={{
          height: '100%',
          '& .simplebar-content': {
            height: '100%'
          }
        }}
      >
        <Box
          sx={{
            display: 'flex',
            flexDirection: 'column',
            height: '100%'
          }}
        >
          <div>
            <Box sx={{ p: 3 }}>
              <NextLink
                href="/dashboard"
                passHref
              >
                <a>
                  <Logo
                    sx={{
                      height: 42,
                      width: 42
                    }}
                  />
                </a>
              </NextLink>
            </Box>
            <Box sx={{ px: 2 }}>
              <Box
                sx={{
                  alignItems: 'center',
                  backgroundColor: 'rgba(255, 255, 255, 0.04)',
                  display: 'flex',
                  justifyContent: 'center',
                  px: 3,
                  py: '11px',
                  borderRadius: 1,
                }}
              >
                <div>
                  <Typography
                    color="inherit"
                    variant="subtitle1"
                  >
                    Corentin Boulanouar
                  </Typography>
                </div>
                {/* <SelectorIcon
                  sx={{
                    color: 'neutral.500',
                    width: 14,
                    height: 14
                  }}
                /> */}
              </Box>
            </Box>
          </div>
          <Divider
            sx={{
              borderColor: '#2D3748', // dark divider
              my: 3
            }}
          />
          <Box sx={{ flexGrow: 1 }}>
            {sections.map((section) => (
              <DashboardSidebarSection
                key={section.title}
                path={router.asPath}
                sx={{
                  mt: 2,
                  '& + &': {
                    mt: 2
                  }
                }}
                {...section}
              />
            ))}
          </Box>
          <Divider
            sx={{
              borderColor: '#2D3748'  // dark divider
            }}
          />
        </Box>
      </Scrollbar>
    </>
  );

  if (lgUp) {
    return (
      <Drawer
        anchor="left"
        open
        PaperProps={{
          sx: {
            backgroundColor: 'neutral.900',
            borderRightColor: 'divider',
            borderRightStyle: 'solid',
            borderRightWidth: (theme) => theme.palette.mode === 'dark' ? 1 : 0,
            color: '#FFFFFF',
            width: 280
          }
        }}
        variant="permanent"
      >
        {content}
      </Drawer>
    );
  }

  return (
    <Drawer
      anchor="left"
      onClose={onClose}
      open={open}
      PaperProps={{
        sx: {
          backgroundColor: 'neutral.900',
          color: '#FFFFFF',
          width: 280
        }
      }}
      sx={{ zIndex: (theme) => theme.zIndex.appBar + 100 }}
      variant="temporary"
    >
      {content}
    </Drawer>
  );
};

DashboardSidebar.propTypes = {
  onClose: PropTypes.func,
  open: PropTypes.bool
};
