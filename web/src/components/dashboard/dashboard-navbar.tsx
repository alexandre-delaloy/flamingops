import { useRef, useState } from 'react';
import type { FC } from 'react';
import PropTypes from 'prop-types';
import {
  AppBar,
  Avatar,
  Box,
  ButtonBase,
  IconButton,
  Toolbar,
} from '@mui/material';
import { styled } from '@mui/material/styles';
import type { AppBarProps } from '@mui/material';
import { Menu as MenuIcon } from '../../icons/menu';
import { AccountPopover } from './account-popover';
import { UserCircle as UserCircleIcon } from '../../icons/user-circle';

interface DashboardNavbarProps extends AppBarProps {
  onOpenSidebar?: () => void;
}

const DashboardNavbarRoot = styled(AppBar)(
  ({ theme }) => ({
    backgroundColor: theme.palette.background.paper,
    ...(
      theme.palette.mode === 'light'
        ? {
          boxShadow: theme.shadows[3]
        }
        : {
          backgroundColor: theme.palette.background.paper,
          borderBottomColor: theme.palette.divider,
          borderBottomStyle: 'solid',
          borderBottomWidth: 1,
          boxShadow: 'none'
        }
    )
  })
);


const AccountButton = () => {
  const anchorRef = useRef<HTMLButtonElement | null>(null);
  const [openPopover, setOpenPopover] = useState<boolean>(false);
  // To get the user from the authContext, you can use
  // `const { user } = useAuth();`
  const user = {
    avatar: '/static/mock-images/avatars/avatar-anika_visser.png',
    name: 'Anika Visser'
  };

  const handleOpenPopover = (): void => {
    setOpenPopover(true);
  };

  const handleClosePopover = (): void => {
    setOpenPopover(false);
  };

  return (
    <>
      <Box
        component={ButtonBase}
        onClick={handleOpenPopover}
        ref={anchorRef}
        sx={{
          alignItems: 'center',
          display: 'flex',
          ml: 2
        }}
      >
        <Avatar
          sx={{
            height: 40,
            width: 40
          }}
          src={user.avatar}
        >
          <UserCircleIcon fontSize="small" />
        </Avatar>
      </Box>
      <AccountPopover
        anchorEl={anchorRef.current}
        onClose={handleClosePopover}
        open={openPopover}
      />
    </>
  );
};

export const DashboardNavbar: FC<DashboardNavbarProps> = (props) => {
  const { onOpenSidebar, ...other } = props;

  return (
    <>
      <DashboardNavbarRoot
        sx={{
          left: {
            lg: 280
          },
          width: {
            lg: 'calc(100% - 280px)'
          }
        }}
        {...other}
      >
        <Toolbar
          disableGutters
          sx={{
            minHeight: 64,
            left: 0,
            px: 2
          }}
        >
          <IconButton
            onClick={onOpenSidebar}
            sx={{
              display: {
                xs: 'inline-flex',
                lg: 'none'
              }
            }}
          >
            <MenuIcon fontSize="small" />
          </IconButton>
          <Box sx={{ flexGrow: 1 }} />
          <AccountButton />
        </Toolbar>
      </DashboardNavbarRoot>
    </>
  );
};

DashboardNavbar.propTypes = {
  onOpenSidebar: PropTypes.func
};
