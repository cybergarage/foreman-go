Name: foreman-go
Version: %{expand:%%(git describe --abbrev=0 --tags)}
Release:	%{_sd_build_id}%{?dist}
Summary: Foreman Go daemon

License: BSD-3-clause
URL: http://github.com/cybergarage/foreman-cc

%{?systemd_requires}
BuildRequires: systemd
BuildRequires: foreman-cc, sqlite-devel, curl-devel, libuuid-devel, python-devel
Requires: sqlite, libuuid, python, curl

Source: %{expand:%%(pwd)}

%description
The Go component of the Foreman system.


%prep
find . -mindepth 1 -delete
cp -af %{SOURCEURL0}/. .
rm -rf src/github.com/cybergarage/foreman-go
mkdir -p src/github.com/cybergarage
ln -s "$PWD" src/github.com/cybergarage/foreman-go
go get ./... || true

%build

%install
export CGO_CXXFLAGS="-static-libstdc++ -static-libgcc"
export CGO_CFLAGS="-static-libgcc"
export go_linker_flags='-linkmode=external "-extldflags=-static-libstdc++ -static-libgcc"'
GOBIN=%{buildroot}/usr/sbin make install
mkdir -p %{buildroot}/etc/foreman
mkdir -p %{buildroot}/etc/foreman/bootstrap
cp debian/foremand.conf %{buildroot}/etc/foreman/foremand.conf
mkdir -p %{buildroot}/%{_unitdir}
cp debian/foremand.service %{buildroot}/%{_unitdir}/foremand.service
mkdir -p %{buildroot}/var/log/foreman

%files
%defattr(755,root,root,755)
/usr/sbin/foremand
/usr/sbin/foremantest
%defattr(644,root,root,755)
/etc/foreman
%{_unitdir}/foremand.service
%dir /var/log/foreman
%dir /etc/foreman/bootstrap

%post
%systemd_post foremand.service

%preun
%systemd_preun foremand.service

%postun
%systemd_postun_with_restart foremand.service

%changelog
* Thu Jan 25 2018 Satoshi Konno <skonno@cybergarage.org>
- Initial packaging
