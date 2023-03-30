import os
import requests
import distro
from datetime import datetime
from threading import Thread

def split_lines(paragraph:str)->list:
    return paragraph.split("\n")

def get_installed_packages_names()->str:
    stream = os.popen("apt list --installed")
    out = stream.read()
    package_full_version_list = out.split("\n")[1:-1]
    package_names = [ pkg.split("/")[0] for pkg in package_full_version_list if "automatic" not in pkg ]
    return package_names


def get_installed_package(lines:list)->str:
    return (lines[1].split(":")[1])[1:]


def get_installed_version(lines:list)->str:
    return get_installed_package(lines).split("-")[0]


def get_candidate_package(lines:list)->str:
    return (lines[2].split(":")[1])[1:]

def get_candidate_version(lines:list)->str:
    return get_candidate_package(lines).split("-")[0]


def get_package_data(pkg,vm):
    stream = os.popen(f"apt-cache policy {pkg}")
    out = stream.read()
    lines = split_lines(out)

    installed_version = get_installed_version(lines)
    candidate_version = get_candidate_version(lines)

    package_info = {
                        "package-name": pkg,
                        "installed-version": installed_version,
                        "condidate-version": candidate_version,
                    }
    vm["installed-packages"].append(package_info)

if __name__== "__main__":

    #get installed pkgs names    
    package_names = get_installed_packages_names()
    
    #get time
    execution_time = datetime.now()
    
    # get external IP
    external_ip = requests.get('https://ident.me').text
    
    # get distro
    dist_info = [distro.name(),distro.version(),distro.codename()]

    vm_pkgs = {
        "ip" : external_ip,
        "distro" : dist_info,
        "date-of-execution" : f"{execution_time}",
        "installed-packages" : []
    }


    threads = []

    for pkg in package_names:
        p = Thread(target=get_package_data, args=[pkg,vm_pkgs])
        threads.append(p)
        p.start()

    for p in threads:
        p.join()

    print(vm_pkgs)

    #TODO send vmpkgs to the backend  
    #TODO check if GIL is slowing the threads  
    