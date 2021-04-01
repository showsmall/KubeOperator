import {Component, OnInit} from '@angular/core';
import {BaseModelDirective} from '../../../../shared/class/BaseModelDirective';
import {Registry} from '../registry';
import {SystemService} from '../../system.service';
import {CommonAlertService} from '../../../../layout/common-alert/common-alert.service';
import {TranslateService} from '@ngx-translate/core';
import {RegistryService} from '../registry.service';
import {System} from '../../system/system';

@Component({
    selector: 'app-registry-list',
    templateUrl: './registry-list.component.html',
    styleUrls: ['./registry-list.component.css']
})
export class RegistryListComponent extends BaseModelDirective<Registry> implements OnInit {
    item: System = new System();

    constructor(private systemService: SystemService, private commonAlertService: CommonAlertService,
                private translateService: TranslateService, private registryService: RegistryService) {
        super(registryService);
    }

    ngOnInit(): void {
        super.ngOnInit();
    }

}
