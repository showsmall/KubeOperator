import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {BaseModelService} from '../../shared/class/BaseModelService';
import {Observable} from 'rxjs';
import {Host, HostSync} from './host';
import {Page} from '../../shared/class/Page';

@Injectable({
    providedIn: 'root'
})
export class HostService extends BaseModelService<Host> {

    baseUrl = '/api/v1/hosts';

    constructor(http: HttpClient) {
        super(http);
    }

    syncList(hosts: HostSync[]): Observable<any> {
        const itemUrl = `${this.baseUrl}/sync/${name}`;
        return this.http.post<HostSync[]>(itemUrl, hosts);
    }

    listByProjectName(projectName: string): Observable<Page<Host>> {
        const itemUrl = `${this.baseUrl}/?projectName=${projectName}`;
        return this.http.get<Page<Host>>(itemUrl);
    }

    upload(file): Observable<any> {
        const itemUrl = `${this.baseUrl}/upload`;
        return this.http.post(itemUrl, file);
    }
}
