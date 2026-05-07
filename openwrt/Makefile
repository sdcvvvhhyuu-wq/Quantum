include $(TOPDIR)/rules.mk

PKG_NAME:=argotunnel
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/argotunnel
  SECTION:=net
  CATEGORY:=Network
  TITLE:=ArgoTunnel anti-censorship
endef

define Package/argotunnel/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) ./files/argotunnel.init $(1)/etc/init.d/argotunnel
endef

$(eval $(call BuildPackage,argotunnel))
