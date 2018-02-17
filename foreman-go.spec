Name: foreman-go
Version: %{expand:%%(git describe --abbrev=0 --tags)}
Release:	1%{?dist}
Summary: Foreman Go daemon

License: BSD-3-clause
URL: http://github.com/cybergarage/foreman-cc

BuildRequires: foreman-cc

Source: %{expand:%%(pwd)}

%description
The Go component of the Foreman system.


%prep
export GOPATH="$PWD"
find . -mindepth 1 -delete
cp -af %{SOURCEURL0}/. .
rm -rf src/github.com/cybergarage/foreman-go
mkdir -p src/github.com/cybergarage
ln -s "$PWD" src/github.com/cybergarage/foreman-go
go get ./... || true
./setup || true

%build

%install
export GOPATH="$PWD"
GOBIN=%{buildroot}/usr/sbin make install
mkdir -p %{buildroot}/etc/foreman
cp debian/foremand.conf %{buildroot}/etc/foreman/foremand.conf

%files
%defattr(755,root,root,755)
/usr/sbin/foremand
/usr/sbin/foremantest
%defattr(644,root,root,755)
/etc/foreman

%changelog
* Thu Jan 25 2018 Satoshi Konno <skonno@cybergarage.org>
- Initial packaging
