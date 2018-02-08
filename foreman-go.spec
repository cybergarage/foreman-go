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
find . -mindepth 1 -delete
cp -af %{SOURCEURL0}/. .

%build
export GOPATH="$PWD"
make build %{?_smp_mflags}


%install
make install DESTDIR=%{buildroot}
cp -a bin/foremand %{buildroot}/usr/sbin

%files
%defattr(755,root,root,755)
/usr/sbin/foremand
#%defattr(755,root,root,755)
#/etc/foreman/foremand.conf

%changelog
* Thu Jan 25 2018 Satoshi Konno <skonno@cybergarage.org>
- Initial packaging
